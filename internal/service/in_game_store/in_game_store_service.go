package in_game_store

import (
	"context"
	"fmt"
	"locgame-mini-server/internal/blockchain"
	"locgame-mini-server/internal/blockchain/contracts"
	"locgame-mini-server/internal/config"
	"locgame-mini-server/internal/service/inventory"
	"locgame-mini-server/internal/service/payments"
	"locgame-mini-server/internal/sessions"
	"locgame-mini-server/internal/store"
	"locgame-mini-server/internal/utils"
	"locgame-mini-server/pkg/dto/base"
	"locgame-mini-server/pkg/dto/cards"
	"locgame-mini-server/pkg/dto/errors"
	"locgame-mini-server/pkg/dto/player"
	"locgame-mini-server/pkg/dto/resources"
	"locgame-mini-server/pkg/log"
	"locgame-mini-server/pkg/pubsub"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/slices"

	protobuf "google.golang.org/protobuf/proto"

	storeDto "locgame-mini-server/pkg/dto/store"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	store      *store.Store
	config     *config.Config
	sessions   *sessions.SessionStore
	blockchain *blockchain.Blockchain
	payments   *payments.Service
	inventory  *inventory.Service

	discounts           map[string]*storeDto.Discount
	nextDiscountsUpdate *base.Timestamp
}

// New creates a new instance of the friends services.
func New(
	config *config.Config,
	sessions *sessions.SessionStore,
	store *store.Store,
	blockchain *blockchain.Blockchain,
	payments *payments.Service,
	inventory *inventory.Service,
) *Service {
	s := new(Service)
	s.config = config
	s.store = store
	s.sessions = sessions
	s.blockchain = blockchain
	s.payments = payments
	s.inventory = inventory

	s.payments.OnPaymentSuccess = s.OnPaymentSuccess

	// pubsub.RegisterPlayerHandler(&CoinsPurchaseCompleteHandler{Sessions: sessions})
	// pubsub.RegisterPlayerHandler(&TokenPurchaseCompleteHandler{Sessions: sessions})
	// pubsub.RegisterPlayerHandler(&PackPurchaseCompleteHandler{})
	// pubsub.RegisterPlayerHandler(&CardUpgradeCompleteHandler{})

	return s
}

func (s *Service) Init() {
	// err := pubsub.Subscribe("update-discounts", &UpdateDiscountsHandler{service: s})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	s.updateDiscounts(false)
}

func (s *Service) updateDiscounts(ignoreErr bool) {
	discounts, err := s.store.Discounts.Get(context.Background())
	if err != nil && !ignoreErr {
		log.Fatal(err)
	}
	now := time.Now().UTC().Unix()
	s.discounts = make(map[string]*storeDto.Discount)
	if len(discounts) > 0 {
		s.nextDiscountsUpdate = &base.Timestamp{Seconds: discounts[0].Duration.EndTime.Seconds}
		for _, discount := range discounts {
			if discount.Active && discount.Duration.EndTime.Seconds > now {
				if discount.Duration.EndTime.Seconds < s.nextDiscountsUpdate.Seconds {
					s.nextDiscountsUpdate.Seconds = discount.Duration.EndTime.Seconds
				}
				s.discounts[discount.Product] = discount
			}
		}
	}
}

func (s *Service) getProducts(
	ctx context.Context,
) (packs []*storeDto.Pack, specialOffers []*storeDto.Pack, coinsPacks []*storeDto.CoinsPack, tokens []*storeDto.Token, err error) {
	productsSold, err := s.store.InGameStore.GetProductsSold(ctx)
	if err != nil {
		return nil, nil, nil, nil, errors.ErrUnexpectedError
	}

	for id, product := range s.config.Products.ProductsByID {
		switch product.Type {
		case storeDto.ProductType_PackOfCards:
			pack := protobuf.Clone(s.config.InGameStore.PackByID[product.Value]).(*storeDto.Pack)
			if pack != nil {
				pack.ProductID = id
				pack.PriceInUSD = product.PriceInUSD
				pack.PriceInLC = product.PriceInLC
				pack.Available = product.Quantity <= 0 || product.Quantity > productsSold[id]
				packs = append(packs, pack)
			}
		case storeDto.ProductType_PackOfCoins:
			coinsPack := s.config.InGameStore.CoinsByID[product.Value]
			if coinsPack != nil {
				coinsPack.ProductID = id
				coinsPack.PriceInUSD = product.PriceInUSD
				coinsPacks = append(coinsPacks, coinsPack)
			}
		case storeDto.ProductType_SpecialOffer:
			pack := s.config.InGameStore.PackByID[product.Value]
			if pack != nil {
				pack.ProductID = id
				pack.PriceInUSD = product.PriceInUSD
				pack.PriceInLC = product.PriceInLC
				pack.Available = product.Quantity <= 0 || product.Quantity > productsSold[id]
				specialOffers = append(specialOffers, pack)
			}
		case storeDto.ProductType_VToken:
			token := protobuf.Clone(s.config.InGameStore.TokensByID[product.Value]).(*storeDto.Token)
			if token != nil {
				token.ProductID = id
				token.PriceInUSD = product.PriceInUSD
				token.Remaining = 0
				checkTokenAvailability(token, productsSold[token.TokenID])
				tokens = append(tokens, token)
			}
		}
	}

	return
}

func (s *Service) getUpgrades() map[string]*storeDto.Upgrades {
	cardsById := s.config.VirtualCards.CardsByID
	upgrades := make(map[string]*storeDto.Upgrades)
	for k, v := range cardsById {
		if v.Upgradable {
			upgrades[k] = &storeDto.Upgrades{
				Options: append([]*cards.Option{}, v.Options...),
			}
		}
	}
	return upgrades
}

func (s *Service) GetData(ctx context.Context) (*storeDto.StoreData, error) {
	packs, specialOffers, coinsPacks, tokens, err := s.getProducts(ctx)
	if err != nil {
		return nil, err
	}
	ethRate, err := s.payments.GetEthRate()
	if err != nil {
		return nil, err
	}
	rate, err := s.payments.GetLoCGRate()
	if err != nil {
		return nil, err
	}
	if s.nextDiscountsUpdate != nil && s.nextDiscountsUpdate.Seconds < time.Now().UTC().Unix() {
		s.updateDiscounts(true)
	}
	upgrades := s.getUpgrades()
	return &storeDto.StoreData{
		LoCGConvertRate: rate,
		EthConvertRate:  ethRate,
		SpecialOffers:   specialOffers,
		Packs:           packs,
		Coins:           coinsPacks,
		Discounts:       s.discounts,
		Upgrades:        upgrades,
		Tokens:          tokens,
	}, nil
}

func getPriceCurrency(product config.Product, paymentMethod storeDto.PaymentMethod) float32 {
	if paymentMethod == storeDto.PaymentMethod_LC {
		if product.PriceInLC > 0 {
			return product.PriceInLC
		} else {
			return product.PriceInUSD / 0.001
		}
	}
	return product.PriceInUSD
}

func (s *Service) getPrice(id string, request *storeDto.OrderRequest, product config.Product) float64 {
	productID := request.ProductID
	paymentMethod := request.PaymentMethod
	productPrice := getPriceCurrency(product, paymentMethod)

	if discount, exists := s.discounts[productID]; exists {
		now := time.Now().UTC().Unix()
		if discount.Duration.StartTime.Seconds < now && discount.Duration.EndTime.Seconds > now {
			switch discount.Type {
			case storeDto.DiscountType_Percentage:
				v := float64(productPrice - (productPrice * (discount.Value / 100.0)))
				return math.Round(v*100) / 100 // Round float to 2 decimal places (round to nearest)
			case storeDto.DiscountType_Fixed:
				v := math.Max(float64(productPrice-discount.Value), 0)
				return math.Round(v*100) / 100
			}
		}
	}

	if product.Type == storeDto.ProductType_VToken {
		token := s.config.InGameStore.TokensByID[product.Value]

		if token == nil {
			return float64(productPrice)
		}

		playerPromoCodeData := s.sessions.Get(id).PlayerData.ActivePromoCode
		var promoCodeConfig *storeDto.TokenPromoCode

		if playerPromoCodeData != nil {
			for _, promo := range token.PromoCodes {
				if promo.ID == playerPromoCodeData.PromoCodeType && promo.Active {
					promoCodeConfig = promo
					break
				}
			}

			if promoCodeConfig != nil {
				isApplicablePaymentMethod := false

				if promoCodeConfig.PaymentMethodsInclude != nil && len(promoCodeConfig.PaymentMethodsInclude) > 0 {
					isIncluded := slices.ContainsFunc(
						promoCodeConfig.PaymentMethodsInclude,
						func(pm int32) bool {
							return storeDto.PaymentMethod(pm) == paymentMethod
						})

					isExcluded := slices.ContainsFunc(
						promoCodeConfig.PaymentMethodsExcept,
						func(pm int32) bool {
							return storeDto.PaymentMethod(pm) == paymentMethod
						})

					isApplicablePaymentMethod = isIncluded && !isExcluded

				} else if promoCodeConfig.PaymentMethodsExcept != nil && len(promoCodeConfig.PaymentMethodsExcept) > 0 {
					isExcluded := slices.ContainsFunc(
						promoCodeConfig.PaymentMethodsExcept,
						func(pm int32) bool {
							return storeDto.PaymentMethod(pm) == paymentMethod
						})

					isApplicablePaymentMethod = !isExcluded
				} else {
					isApplicablePaymentMethod = true
				}

				if isApplicablePaymentMethod {
					var bonus float32

					if playerPromoCodeData.IsOwner {
						bonus = promoCodeConfig.Bonus
					} else {
						bonus = promoCodeConfig.ReferralBonus
					}

					productPrice = productPrice / (1 + bonus)
				}
			}
		}

		if paymentMethod == storeDto.PaymentMethod_LOCG || paymentMethod == storeDto.PaymentMethod_LOCGBase {
			now := time.Now().UTC().Unix()
			offerApplied := false
			for _, offer := range token.Offers {
				if offer.SaleStart < now && offer.SaleEnd > now {
					if offer.LOCGBonusOverride > 0 {
						productPrice = productPrice / (1 + offer.LOCGBonusOverride)
						offerApplied = true
						break
					}
				}
			}
			if !offerApplied && token.LOCGBonus > 0 {
				productPrice = productPrice / (1 + token.LOCGBonus)
			}
		}
	}
	return float64(productPrice)
}

func (s *Service) CreateOrder(
	ctx context.Context, id string,
	request *storeDto.OrderRequest,
) (*storeDto.OrderResponse, error) {
	if request.Quantity < 1 {
		return nil, errors.ErrInvalidQuantity
	}

	product, ok := s.config.Products.ProductsByID[request.ProductID]
	if !ok {
		return nil, errors.ErrProductNotFound
	}
	var token *storeDto.Token
	var productId string

	if product.Type == storeDto.ProductType_VToken {
		token = s.config.InGameStore.TokensByID[product.Value]
		productId = token.TokenID
	} else {
		productId = request.ProductID
	}

	count, err := s.store.InGameStore.GetProductSold(ctx, productId)
	if err != nil {
		return nil, err
	}

	if product.Quantity > 0 && count >= product.Quantity {
		return nil, errors.ErrProductIsSoldOut
	}

	if request.PaymentMethod == storeDto.PaymentMethod_LC &&
		(product.Type == storeDto.ProductType_PackOfCoins || product.Type == storeDto.ProductType_VToken) {
		return nil, errors.ErrUnexpectedError
	}

	session := s.sessions.Get(id)

	if product.Type == storeDto.ProductType_SpecialOffer {
		err := s.validateSpecialOffer(session, &product)
		if err != nil {
			return nil, err
		}
	}

	if (s.config.Environment != config.Development && s.config.Environment != config.Staging && s.config.Environment != config.Production) &&
		(product.Type == storeDto.ProductType_PackOfCards || product.Type == storeDto.ProductType_SpecialOffer) {
		return nil, errors.ErrOperationIsProhibited
	}

	var price *big.Int

	switch request.PaymentMethod {
	case storeDto.PaymentMethod_USDC,
		storeDto.PaymentMethod_USDCBase,
		storeDto.PaymentMethod_USDT:
		price = payments.ToWei(s.getPrice(id, request, product), 6)
	case storeDto.PaymentMethod_LOCG,
		storeDto.PaymentMethod_LOCGBase:
		rate, err := s.payments.GetLoCGRate()
		if err != nil {
			return nil, err
		}
		price = payments.ToWei(s.getPrice(id, request, product)*rate.Price, 18)
	case storeDto.PaymentMethod_LC:
		price = big.NewInt(int64(s.getPrice(id, request, product)))
	case storeDto.PaymentMethod_ETH,
		storeDto.PaymentMethod_ETHBase:
		rate, err := s.payments.GetEthRate()
		if err != nil {
			return nil, err
		}
		price = payments.ToWei(s.getPrice(id, request, product)*rate.Price, 18)
	}

	var maxQuantity int64 = 1
	var minQuantity int32 = 1

	switch product.Type {
	case storeDto.ProductType_PackOfCards:
		pack := s.config.InGameStore.PackByID[product.Value]
		if pack == nil {
			return nil, errors.ErrProductNotFound
		}
		maxQuantity = int64(pack.MaxPurchase)
	case storeDto.ProductType_PackOfCoins:
		pack := s.config.InGameStore.CoinsByID[product.Value]
		if pack == nil {
			return nil, errors.ErrProductNotFound
		}
		maxQuantity = 1
	case storeDto.ProductType_VToken:
		if token == nil || !token.Available {
			return nil, errors.ErrProductNotFound
		}
		tokensAvailable := token.MaxSupply - int64(count)
		if tokensAvailable < 1 {
			return nil, errors.ErrProductIsSoldOut
		}
		tokensRequired := request.Quantity * int64(token.QtyPerUnit)
		if int64(tokensRequired) > tokensAvailable {
			// calculate max quantity
			maxQuantity = int64(math.Floor(float64(tokensAvailable) / float64(token.QtyPerUnit)))
			if maxQuantity < 1 {
				return nil, errors.ErrNotEnoughResources
			}
		} else {
			maxQuantity = token.MaxPurchase
		}
		if token.MinPurchase > 1 {
			minQuantity = token.MinPurchase
		}
	}

	if request.Quantity > maxQuantity {
		request.Quantity = maxQuantity
	}
	if request.Quantity < int64(minQuantity) {
		return nil, errors.ErrInvalidQuantity
	}
	if price != nil {
		price = price.Mul(price, big.NewInt(request.Quantity))
	}
	priceUSD := s.getPrice(id, request, product) * float64(request.Quantity)

	var total string
	if price != nil {
		total = price.String()
	} else {
		total = fmt.Sprintf("%.2f", priceUSD)
	}

	order := &storeDto.Order{
		BuyerID:       session.PlayerData.ID,
		ProductID:     request.ProductID,
		PaymentMethod: request.PaymentMethod,
		ProductType:   product.Type,
		Quantity:      request.Quantity,
		Status:        storeDto.OrderStatus_WaitingForPayment,
		Price:         total,
		CreatedAt:     &base.Timestamp{Seconds: time.Now().UTC().Unix()},
	}

	orderID, err := s.store.Orders.Create(ctx, order)
	order.ID = &base.ObjectID{Value: orderID}
	if err == nil {
		switch request.PaymentMethod {
		case storeDto.PaymentMethod_USDC:
			b, err := blockchain.NewBalanceChecker(
				s.config.Blockchain.RpcAddresses.Ethereum,
				s.config.Blockchain.Contracts.USDC,
			)
			if err != nil {
				log.Debug(err)
				return nil, errors.ErrNotEnoughResources
			}

			session := s.sessions.Get(id)

			err = b.IsTransferable(
				price,
				common.HexToAddress(session.PlayerData.ActiveWallet),
				common.HexToAddress(s.config.Blockchain.PaymentRecipients.USDC),
			)
			if err != nil {
				log.Debug(err)
				return nil, errors.ErrNotEnoughResources
			}

			abi, _ := contracts.ERC20MetaData.GetAbi()
			pack, _ := abi.Pack(
				"transfer",
				common.HexToAddress(s.config.Blockchain.PaymentRecipients.USDC),
				price,
			)
			return &storeDto.OrderResponse{
				Order:    order,
				To:       s.config.Blockchain.Contracts.USDC,
				CallData: "0x" + common.Bytes2Hex(pack),
				ChainID:  s.config.Blockchain.ChainIds.Ethereum,
			}, err
		case storeDto.PaymentMethod_USDT:
			b, err := blockchain.NewBalanceChecker(
				s.config.Blockchain.RpcAddresses.Ethereum,
				s.config.Blockchain.Contracts.USDT,
			)
			if err != nil {
				log.Debug(err)
				return nil, errors.ErrNotEnoughResources
			}

			session := s.sessions.Get(id)

			err = b.IsTransferable(
				price,
				common.HexToAddress(session.PlayerData.ActiveWallet),
				common.HexToAddress(s.config.Blockchain.PaymentRecipients.USDT),
			)
			if err != nil {
				log.Debug(err)
				return nil, errors.ErrNotEnoughResources
			}

			abi, _ := contracts.ERC20MetaData.GetAbi()
			pack, _ := abi.Pack(
				"transfer",
				common.HexToAddress(s.config.Blockchain.PaymentRecipients.USDT),
				price,
			)
			return &storeDto.OrderResponse{
				Order:    order,
				To:       s.config.Blockchain.Contracts.USDT,
				CallData: "0x" + common.Bytes2Hex(pack),
				ChainID:  s.config.Blockchain.ChainIds.Ethereum,
			}, err
		case storeDto.PaymentMethod_LOCG:
			b, err := blockchain.NewBalanceChecker(
				s.config.Blockchain.RpcAddresses.Ethereum,
				s.config.Blockchain.Contracts.LOCG,
			)
			if err != nil {
				log.Debug(err)
				return nil, errors.ErrNotEnoughResources
			}

			session := s.sessions.Get(id)

			err = b.IsTransferable(
				price,
				common.HexToAddress(session.PlayerData.ActiveWallet),
				common.HexToAddress(s.config.Blockchain.PaymentRecipients.LOCG),
			)
			if err != nil {
				log.Debug(err)
				return nil, errors.ErrNotEnoughResources
			}

			abi, _ := contracts.ERC20MetaData.GetAbi()
			pack, _ := abi.Pack(
				"transfer",
				common.HexToAddress(s.config.Blockchain.PaymentRecipients.LOCG),
				price,
			)
			return &storeDto.OrderResponse{
				Order:    order,
				To:       s.config.Blockchain.Contracts.LOCG,
				CallData: "0x" + common.Bytes2Hex(pack),
				ChainID:  s.config.Blockchain.ChainIds.Ethereum,
			}, err
		case storeDto.PaymentMethod_LOCGBase:
			b, err := blockchain.NewBalanceChecker(
				s.config.Blockchain.RpcAddresses.Base,
				s.config.Blockchain.Contracts.BaseLOCG,
			)
			if err != nil {
				log.Debug(err)
				return nil, errors.ErrNotEnoughResources
			}

			session := s.sessions.Get(id)

			err = b.IsTransferable(
				price,
				common.HexToAddress(session.PlayerData.ActiveWallet),
				common.HexToAddress(s.config.Blockchain.PaymentRecipients.LOCG),
			)
			if err != nil {
				log.Debug(err)
				return nil, errors.ErrNotEnoughResources
			}

			abi, _ := contracts.LOCGBridgedMetaData.GetAbi()
			pack, _ := abi.Pack(
				"transfer",
				common.HexToAddress(s.config.Blockchain.PaymentRecipients.LOCG),
				price,
			)
			return &storeDto.OrderResponse{
				Order:    order,
				To:       s.config.Blockchain.Contracts.BaseLOCG,
				CallData: "0x" + common.Bytes2Hex(pack),
				ChainID:  s.config.Blockchain.ChainIds.Base,
			}, err
		case storeDto.PaymentMethod_USDCBase:
			b, err := blockchain.NewBalanceChecker(
				s.config.Blockchain.RpcAddresses.Base,
				s.config.Blockchain.Contracts.BaseUSDC,
			)
			if err != nil {
				log.Debug(err)
				return nil, errors.ErrNotEnoughResources
			}

			session := s.sessions.Get(id)

			err = b.IsTransferable(
				price,
				common.HexToAddress(session.PlayerData.ActiveWallet),
				common.HexToAddress(s.config.Blockchain.PaymentRecipients.USDC),
			)
			if err != nil {
				log.Debug(err)
				return nil, errors.ErrNotEnoughResources
			}

			abi, _ := contracts.LOCGBridgedMetaData.GetAbi()
			pack, _ := abi.Pack(
				"transfer",
				common.HexToAddress(s.config.Blockchain.PaymentRecipients.LOCG),
				price,
			)
			return &storeDto.OrderResponse{
				Order:    order,
				To:       s.config.Blockchain.Contracts.BaseUSDC,
				CallData: "0x" + common.Bytes2Hex(pack),
				ChainID:  s.config.Blockchain.ChainIds.Base,
			}, err
		case storeDto.PaymentMethod_LC:
			cost := &resources.ResourceAdjustment{
				ResourceID: 1,
				Quantity:   int32(price.Int64()),
			}
			if !s.inventory.IsEnough(session, cost) {
				return nil, errors.ErrNotEnoughResources
			}
			adjustments, err := s.inventory.InverseAdjust(session, "store", cost)
			if err != nil {
				return nil, err
			}

			order.Status = storeDto.OrderStatus_PaymentReceived
			err = s.store.Orders.Update(ctx, order)
			if err != nil {
				log.Error("Unable to update order in database:", err)
			}

			s.OnPaymentSuccess(ctx, order)

			return &storeDto.OrderResponse{
				Order:       order,
				Adjustments: adjustments,
			}, err
		case storeDto.PaymentMethod_PayPal:
			price, _ := strconv.ParseFloat(total, 32)
			if price > 201 {
				return nil, errors.ErrNotAuthorized
			}
			url, err := s.payments.CreatePaypalPayment(order)
			if err != nil {
				return nil, err
			}
			return &storeDto.OrderResponse{
				Order:    order,
				CallData: url,
			}, err
		case storeDto.PaymentMethod_ETH:
			b, err := blockchain.NewBalanceChecker(
				s.config.Blockchain.RpcAddresses.Ethereum,
				"",
			)
			if err != nil {
				log.Debug(err)
				return nil, errors.ErrNotEnoughResources
			}

			session := s.sessions.Get(id)

			err = b.IsTransferable(
				price,
				common.HexToAddress(session.PlayerData.ActiveWallet),
				common.HexToAddress(s.config.Blockchain.PaymentRecipients.LOCG),
			)
			if err != nil {
				log.Debug(err)
				return nil, errors.ErrNotEnoughResources
			}

			return &storeDto.OrderResponse{
				Order:    order,
				To:       s.config.Blockchain.PaymentRecipients.LOCG,
				CallData: "0x",
				Value:    price.String(),
				ChainID:  s.config.Blockchain.ChainIds.Ethereum,
			}, err
		case storeDto.PaymentMethod_ETHBase:
			b, err := blockchain.NewBalanceChecker(
				s.config.Blockchain.RpcAddresses.Base,
				"",
			)
			if err != nil {
				log.Debug(err)
				return nil, errors.ErrNotEnoughResources
			}

			session := s.sessions.Get(id)

			err = b.IsTransferable(
				price,
				common.HexToAddress(session.PlayerData.ActiveWallet),
				common.HexToAddress(s.config.Blockchain.PaymentRecipients.LOCG),
			)
			if err != nil {
				log.Debug(err)
				return nil, errors.ErrNotEnoughResources
			}

			return &storeDto.OrderResponse{
				Order:    order,
				To:       s.config.Blockchain.PaymentRecipients.LOCG,
				CallData: "0x",
				Value:    price.String(),
				ChainID:  s.config.Blockchain.ChainIds.Base,
			}, err
		}
	}

	return nil, err
}

func (s *Service) CreateUpgradeOrder(
	ctx context.Context,
	id string,
	request *storeDto.UpgradeRequest,
) (*storeDto.OrderResponse, error) {
	session := s.sessions.Get(id)
	if session == nil {
		return nil, errors.ErrSessionNotFound
	}

	player, err := s.store.Players.GetPlayerDataByID(ctx, session.AccountID)
	if err != nil {
		log.Debug("Error fetching User: ", err)
		return nil, err
	}

	if request.PaymentMethod != storeDto.PaymentMethod_LC {
		return nil, errors.ErrInvalidToken
	}

	card, err := s.config.VirtualCards.GetCardById(request.CardId)
	if err != nil || !card.Upgradable {
		if err == nil {
			err = errors.ErrOperationIsProhibited
		}
		return nil, err
	}

	// NOTE: Check Virtual Cards for existing card to upgrade. will need refactoring if other items are made upgradable.

	if !utils.Includes(player.VirtualCards, card.ArchetypeID) {
		return nil, errors.ErrInvalidCard
	}

	var upgradeId string
	var price float32

	for _, option := range card.Options {
		if option.ItemID == request.UpgradeId {
			upgradeId = request.UpgradeId
			price = option.Price
		}
	}
	log.Debug("Upgrade Id: ", upgradeId)
	if upgradeId == "" {
		return nil, errors.ErrInvalidCard
	}

	order := &storeDto.Order{
		BuyerID: player.ID,
		// ProductID set to OriginalID-UpgradeID
		ProductID:     fmt.Sprintf("%s|%s", request.CardId, request.UpgradeId),
		PaymentMethod: request.PaymentMethod,
		ProductType:   storeDto.ProductType_CardUpgrade,
		Quantity:      1,
		Status:        storeDto.OrderStatus_WaitingForPayment,
		Price:         fmt.Sprintf("%f", price),
		CreatedAt:     &base.Timestamp{Seconds: time.Now().UTC().Unix()},
	}

	orderID, err := s.store.Orders.Create(ctx, order)
	if err != nil {
		return nil, err
	}
	order.ID = &base.ObjectID{Value: orderID}
	cost := &resources.ResourceAdjustment{
		ResourceID: 1,
		Quantity:   int32(price),
	}

	adjustments, err := s.inventory.InverseAdjust(session, "upgrade", cost)
	if err != nil {
		return nil, err
	}
	order.Status = storeDto.OrderStatus_PaymentReceived
	err = s.store.Orders.Update(ctx, order)
	if err != nil {
		log.Error("Unable to update order in database:", err)
	}
	s.OnPaymentSuccess(ctx, order)
	return &storeDto.OrderResponse{
		Order:       order,
		Adjustments: adjustments,
	}, nil
}

func (s *Service) validateSpecialOffer(session *sessions.Session, product *config.Product) error {
	switch product.Value {
	case "starter_pack":
		if session.PlayerData.PlayerStoreData.PurchasedProducts[product.Value] > 0 {
			return errors.ErrProductNotAvailable
		}
	case "story_mode_pack":
		if session.PlayerData.PlayerStoreData.PurchasedProducts[product.Value] > 0 ||
			session.PlayerData.StoryMode.LastUnlockedLevel < int32(len(s.config.StoryMode)) {
			return errors.ErrProductNotAvailable
		}
	}
	return nil
}

func (s *Service) SetPaymentReceipt(ctx context.Context, in *storeDto.Receipt) error {
	order, err := s.store.Orders.Get(ctx, in.OrderID)
	if err != nil {
		log.Error("Failed to get order:", in.OrderID, "Error:", err)
		return err
	}

	if order.Status != storeDto.OrderStatus_WaitingForPayment {
		return errors.ErrOrderAlreadyPaid
	}

	order.PaymentHash = in.TransactionHash
	return s.store.Orders.Update(ctx, order)
}

func (s *Service) OnPaymentSuccess(ctx context.Context, order *storeDto.Order) {
	productType := order.ProductType
	switch productType {
	case storeDto.ProductType_PackOfCards, storeDto.ProductType_SpecialOffer:
		_ = s.store.InGameStore.IncrementPurchases(
			context.Background(),
			order.BuyerID,
			order.ProductID,
			order.Quantity,
		)
		_ = s.store.InGameStore.IncrementSold(context.Background(), order.ProductID, order.Quantity)
		err := pubsub.SendEvent(
			&storeDto.MintJobRequest{ID: order.ID, MintType: storeDto.MintType_MintOrder},
		)
		if err != nil {
			log.Error("Failed to send mint event:", order.ID, "Error:", err)
		}
	case storeDto.ProductType_PackOfCoins:
		product := s.config.Products.ProductsByID[order.ProductID]
		_ = s.store.InGameStore.IncrementPurchases(
			context.Background(),
			order.BuyerID,
			order.ProductID,
			order.Quantity,
		)
		err := s.onSuccessfulPurchaseOfCoins(order, product)
		if err != nil {
			log.Error("Failed to adjust coins:", err)
		}
	case storeDto.ProductType_CardUpgrade:
		accountId, _ := primitive.ObjectIDFromHex(order.BuyerID.Value)
		session, _ := s.sessions.TryGetByAccountID(accountId)
		ids := strings.Split(order.ProductID, "|")
		cardType := ids[1][:3]
		if cardType == "999" {
			err := s.updateVirtualCard(ctx, session.SessionID, ids, order)
			if err != nil {
				log.Error("Failed to update virtual card:", err)
			}
		} else {
			// TODO: Update w/ Mint
			err := s.updateCardAndMint(ctx, session.SessionID, ids, order)
			if err != nil {
				log.Error("Failed to update card and mint:", err)
			}
		}
	case storeDto.ProductType_VToken:

		_ = s.store.InGameStore.IncrementPurchases(
			context.Background(),
			order.BuyerID,
			order.ProductID,
			order.Quantity,
		)
		// Update Packs of Tokens Sold
		_ = s.store.InGameStore.IncrementSold(context.Background(), order.ProductID, order.Quantity)

		// Update Total Tokens Sold by Token ID
		product := s.config.Products.ProductsByID[order.ProductID]
		token := s.config.InGameStore.TokensByID[product.Value]
		qty := order.Quantity * int64(token.QtyPerUnit)
		_ = s.store.InGameStore.IncrementSold(context.Background(), token.TokenID, qty)

		err := s.onSuccessfulPurchaseOfTokens(order)
		if err != nil {
			log.Error("Failed to adjust coins:", err)
		}

	}
}

func (s *Service) onSuccessfulPurchaseOfCoins(order *storeDto.Order, product config.Product) error {
	coinsData := s.config.InGameStore.CoinsByID[product.Value]
	adjustment := &resources.ResourceAdjustment{
		ResourceID: 1,
		Quantity:   coinsData.Count,
	}

	accountID, _ := primitive.ObjectIDFromHex(order.BuyerID.Value)
	err := s.store.Inventory.IncrementResources(
		context.Background(),
		accountID,
		[]*resources.ResourceAdjustment{adjustment},
		"coin_purchased",
	)
	if err != nil {
		return err
	}

	// _ = pubsub.SendToPlayer(order.BuyerID.Value, &storeDto.CoinsPurchaseResult{
	// 	Coins:   adjustment,
	// 	OrderID: order.ID,
	// })

	order.Status = storeDto.OrderStatus_Completed
	err = s.store.Orders.Update(context.Background(), order)
	if err != nil {
		log.Error("Unable to update order in database:", err)
	}

	return nil
}

func (s *Service) onSuccessfulPurchaseOfTokens(order *storeDto.Order) error {
	product := s.config.Products.ProductsByID[order.ProductID]
	token := s.config.InGameStore.TokensByID[product.Value]
	qty := order.Quantity * int64(token.QtyPerUnit)

	adjustment := &resources.ResourceAdjustment{
		ResourceID: 3,
		Quantity:   int32(qty),
	}

	accountID, _ := primitive.ObjectIDFromHex(order.BuyerID.Value)
	err := s.store.Inventory.IncrementResources(
		context.Background(),
		accountID,
		[]*resources.ResourceAdjustment{adjustment},
		"token_purchased",
	)
	if err != nil {
		return err
	}

	// _ = pubsub.SendToPlayer(order.BuyerID.Value, &storeDto.TokenPurchaseResult{
	// 	Tokens:  adjustment,
	// 	OrderID: order.ID,
	// })

	order.Status = storeDto.OrderStatus_Completed
	err = s.store.Orders.Update(context.Background(), order)
	if err != nil {
		log.Error("Unable to update order in database:", err)
	}

	return nil
}

func (s *Service) GetUnopenedPacks(ctx context.Context, id string) ([]*storeDto.Order, error) {
	session := s.sessions.Get(id)
	return s.store.Orders.GetUnopenedPacks(ctx, session.PlayerData.ID)
}

func (s *Service) OpenPack(
	ctx context.Context,
	in *base.ObjectID,
) (*storeDto.OpenPackResponse, error) {
	order, err := s.store.Orders.Get(ctx, in.Value)
	if err != nil {
		return nil, err
	}

	if order.Status != storeDto.OrderStatus_Completed {
		return nil, errors.ErrInvalidOrderStatus
	}
	order.Status = storeDto.OrderStatus_Opened
	_ = s.store.Orders.Update(ctx, order)

	return &storeDto.OpenPackResponse{
		Cards: order.Cards,
	}, nil
}

func (s *Service) updateVirtualCard(
	ctx context.Context,
	sessionId string,
	ids []string,
	order *storeDto.Order,
) error {
	session := s.sessions.Get(sessionId)
	playerData, err := s.store.Players.GetPlayerDataByID(ctx, session.AccountID)
	if err != nil {
		log.Debug("Unable to get player")
		return err
	}
	utils.ReplaceValue(&playerData.VirtualCards, ids[0], ids[1])
	err = s.store.Players.SetData(ctx, session.AccountID, playerData)
	if err != nil {
		log.Errorf("Error updating player data: %v", err)
	}
	session.GetData(sessionId)
	return nil
}

func (s *Service) updateCardAndMint(
	ctx context.Context,
	sessionId string,
	ids []string,
	order *storeDto.Order,
) error {
	session := s.sessions.Get(sessionId)
	playerData, err := s.store.Players.GetPlayerDataByID(ctx, session.AccountID)
	if err != nil {
		log.Debug("Unable to get player")
		return err
	}
	utils.RemoveValue(&playerData.VirtualCards, ids[0])
	err = s.store.Players.SetData(ctx, session.AccountID, playerData)
	if err != nil {
		log.Errorf("Error updating player data: %v", err)
	}
	err = pubsub.SendEvent(
		&storeDto.MintJobRequest{ID: order.ID, MintType: storeDto.MintType_MintUpgrade},
	)
	if err != nil {
		log.Error("Failed to send mint event:", order.ID, "Error:", err)
	}
	session.GetData(sessionId)
	return nil
}

func checkTokenAvailability(token *storeDto.Token, soldQty int64) {
	currentTimeUTC := time.Now().UTC()
	fake := int64(0)
	if token.AdditionalAccountedQuantity > 0 {
		fake = token.AdditionalAccountedQuantity
	}
	if currentTimeUTC.Unix() < token.SaleStart || currentTimeUTC.Unix() > token.SaleEnd {
		token.Available = false
	}
	if token.MaxSupply > 0 && int64(soldQty) >= token.MaxSupply {
		token.Available = false
	}
	token.Remaining = token.MaxSupply - soldQty - fake
}

func (s *Service) SubmitPromoCode(ctx context.Context, id string, in *storeDto.PromoCodeSubmitRequest) (*storeDto.PromoCodeSubmitResponse, error) {
	session := s.sessions.Get(id)
	promoCodeData, err := s.store.InGameStore.GetPromoCodeData(ctx, in.PromoCode, session.AccountID)
	if err != nil {
		return &storeDto.PromoCodeSubmitResponse{
				Success: false, PromoCodeTypeId: "", Message: "Error getting promo code",
			},
			nil
	}

	promoCodeTypeId := promoCodeData.PromoCodeTypeId

	if promoCodeData == nil {
		return &storeDto.PromoCodeSubmitResponse{
				Success: false, PromoCodeTypeId: "", Message: "Promo code not found",
			},
			nil
	}

	playerPromoCodeData := &player.PromoCodeData{
		PromoCode:     in.PromoCode,
		PromoCodeType: promoCodeTypeId,
		SubmissionDate: &base.Timestamp{
			Seconds: time.Now().UTC().Unix(),
		},
		IsOwner: promoCodeData.IsOwner,
	}

	err = s.store.Players.SetData(ctx, session.AccountID, &player.PlayerData{ActivePromoCode: playerPromoCodeData})
	if err != nil {
		return &storeDto.PromoCodeSubmitResponse{
				Success: false, PromoCodeTypeId: "", Message: "Failed to update player data 1",
			},
			nil
	}

	err = s.sessions.Get(id).PlayerData.Update(ctx)
	if err != nil {
		return &storeDto.PromoCodeSubmitResponse{
				Success: false, PromoCodeTypeId: "", Message: "Failed to update player data 2",
			},
			nil
	}

	return &storeDto.PromoCodeSubmitResponse{
			Success: true, PromoCodeTypeId: promoCodeTypeId,
		},
		nil
}
