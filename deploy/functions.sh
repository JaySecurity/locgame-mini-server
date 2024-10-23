export ENVIRONMENT
export VERSION
export SERVICE_PORT
export REST_PORT
export REDIS_PORT
export NATS_PORT


function get_available_port() {
    /usr/bin/python3 -c 'import socket; s=socket.socket(); s.bind(("", 0)); print(s.getsockname()[1]); s.close()'
}

# function locg_deploy() {
#     ENVIRONMENT=$1
#     VERSION=$2


#     # SERVICE_PORT=$(get_available_port)
#     SERVICE_PORT=56377
#     # MONGODB_PORT=$(get_available_port)
#     REDIS_PORT=$(get_available_port)
#     NATS_PORT=$(get_available_port)

#     # gcloud dns record-sets create "${ENVIRONMENT}".locg.furylion.dev --rrdatas=35.241.200.151 --ttl=3600 --type=A --zone=furylion-dev --project=furylion-209714

#     echo "Sleeping for 10 seconds…"
#     sleep 10

#     # shellcheck disable=SC2143
#     if [[ $(docker ps -a | grep locg_"${ENVIRONMENT}"_service_1) ]]; then
#         echo "Environment already exists. Deploying with saving the last ports."
#         SERVICE_PORT=$(docker inspect --format "{{index (index (index .NetworkSettings.Ports \"8080/tcp\") 0) \"HostPort\"}}" locg_"${ENVIRONMENT}"_service_1)
#         # MONGODB_PORT=$(docker inspect --format "{{index (index (index .NetworkSettings.Ports \"27017/tcp\") 0) \"HostPort\"}}" locg_"${ENVIRONMENT}"_mongodb_1)
#         REDIS_PORT=$(docker inspect --format "{{index (index (index .NetworkSettings.Ports \"6379/tcp\") 0) \"HostPort\"}}" locg_"${ENVIRONMENT}"_redis_1)
#         NATS_PORT=$(docker inspect --format "{{index (index (index .NetworkSettings.Ports \"4222/tcp\") 0) \"HostPort\"}}" locg_"${ENVIRONMENT}"_nats_1)

#         # if [[ "$SERVICE_PORT" =~ ^[0-9]+$ ]]; then
#         #   echo "HTTP port: OK"
#         # else
#         #   echo "Unable get HTTP port: $SERVICE_PORT"
#         #   SERVICE_PORT=$(get_available_port)
#         # fi

#         # if [[ "$MONGODB_PORT" =~ ^[0-9]+$ ]]; then
#         #   echo "MongoDB port: OK"
#         # else
#         #   echo "Unable get MongoDB port: $MONGODB_PORT"
#         #   MONGODB_PORT=$(get_available_port)
#         # fi

#         if [[ "$REDIS_PORT" =~ ^[0-9]+$ ]]; then
#             echo "Redis port: OK"
#         else
#             echo "Unable get Redis port: $REDIS_PORT"
#             REDIS_PORT=$(get_available_port)
#         fi

#         if [[ "$NATS_PORT" =~ ^[0-9]+$ ]]; then
#             echo "NATS port: OK"
#         else
#             echo "Unable get NATS port: NATS_PORT"
#             NATS_PORT=$(get_available_port)
#         fi
#     fi

#     # Set nginx config with new port env variables
#     # cd
#     # envsubst < /etc/nginx/conf.d/default.conf.template > default.conf
#     # sudo chown root:root default.conf
#     # sudo mv default.conf /etc/nginx/conf.d/default.conf
#     # sudo nginx -s reload
#     # sudo systemctl restart nginx.service


#     echo "Creating internal network..."
#     docker network create internal_${ENVIRONMENT}


#     echo "Pulling..."
#     docker-compose -f ~/locg/deploy/docker-compose.yaml --project-name locg_"$ENVIRONMENT" pull

#     echo "Restarting..."
#     docker-compose -f ~/locg/deploy/docker-compose.yaml --project-name locg_"$ENVIRONMENT" up -d

#     echo "Done"
# }

function locg_delete_environment() {
    echo "Deleting Existing Environment... "
    ENVIRONMENT=$1
    
    SERVICE_PORT=$(docker inspect --format "{{index (index (index .NetworkSettings.Ports \"8080/tcp\") 0) \"HostPort\"}}" locg_"${ENVIRONMENT}"_service_1)
    REST_PORT=$(docker inspect --format "{{index (index (index .NetworkSettings.Ports \"27017/tcp\") 0) \"HostPort\"}}" locg_"${ENVIRONMENT}"_rest-api_1)
    REDIS_PORT=$(docker inspect --format "{{index (index (index .NetworkSettings.Ports \"6379/tcp\") 0) \"HostPort\"}}" locg_"${ENVIRONMENT}"_redis_1)
    NATS_PORT=$(docker inspect --format "{{index (index (index .NetworkSettings.Ports \"4222/tcp\") 0) \"HostPort\"}}" locg_"${ENVIRONMENT}"_nats_1)
    
    docker-compose -f ~/locg/deploy/docker-compose.yaml --project-name locg_"$ENVIRONMENT" rm -f -s
    
    docker network prune -f
    docker image prune -a -f
    
    # gcloud dns record-sets delete "${ENVIRONMENT}".locg.furylion.dev --type=A --zone=furylion-dev --project=furylion-209714
    echo "Done. Environment removed."
}

function locg_deploy() {
    echo "Deploying New Environment: $ENVIRONMENT"
    
    BUILD_ENVIRONMENT=$1
    VERSION=$2
    
    # if [[ "$BUILD_ENVIRONMENT" == "development" ]]; then
    #     SERVICE_PORT=53677
    #     elif [[ "$BUILD_ENVIRONMENT" == "special" ]]; then
    #     SERVICE_PORT=53678
    # else
    #     SERVICE_PORT=53679
    # fi
    
    SERVICE_PORT=$(get_available_port)
    REST_PORT=$(get_available_port)
    REDIS_PORT=$(get_available_port)
    NATS_PORT=$(get_available_port)
    
    # gcloud dns record-sets create "${BUILD_ENVIRONMENT}".locg.furylion.dev --rrdatas=35.241.200.151 --ttl=3600 --type=A --zone=furylion-dev --project=furylion-209714
    
    echo "Sleeping for 10 seconds…"
    sleep 10
    
    # shellcheck disable=SC2143
    if [[ $(docker ps -a | grep locg_"${BUILD_ENVIRONMENT}"_service_1) ]]; then
        echo "BUILD_Environment already exists. Deploying with saving the last ports."
        SERVICE_PORT=$(docker inspect --format "{{index (index (index .NetworkSettings.Ports \"8080/tcp\") 0) \"HostPort\"}}" locg_"${BUILD_ENVIRONMENT}"_service_1)
        REST_PORT=$(docker inspect --format "{{index (index (index .NetworkSettings.Ports \"27017/tcp\") 0) \"HostPort\"}}" locg_"${BUILD_ENVIRONMENT}"_rest-api_1)
        REDIS_PORT=$(docker inspect --format "{{index (index (index .NetworkSettings.Ports \"6379/tcp\") 0) \"HostPort\"}}" locg_"${BUILD_ENVIRONMENT}"_redis_1)
        NATS_PORT=$(docker inspect --format "{{index (index (index .NetworkSettings.Ports \"4222/tcp\") 0) \"HostPort\"}}" locg_"${BUILD_ENVIRONMENT}"_nats_1)
        
        if [[ "$SERVICE_PORT" =~ ^[0-9]+$ ]]; then
            echo "HTTP port: OK"
        else
            echo "Unable get HTTP port: $SERVICE_PORT"
            SERVICE_PORT=$(get_available_port)
        fi
        
        if [[ "$REST_PORT" =~ ^[0-9]+$ ]]; then
            echo "MongoDB port: OK"
        else
            echo "Unable get Rest API port: $REST_PORT"
            REST_PORT=$(get_available_port)
        fi
        
        if [[ "$REDIS_PORT" =~ ^[0-9]+$ ]]; then
            echo "Redis port: OK"
        else
            echo "Unable get Redis port: $REDIS_PORT"
            REDIS_PORT=$(get_available_port)
        fi
        
        if [[ "$NATS_PORT" =~ ^[0-9]+$ ]]; then
            echo "NATS port: OK"
        else
            echo "Unable get NATS port: NATS_PORT"
            NATS_PORT=$(get_available_port)
        fi
    fi
    
    # Set nginx config with new port env variables
    # cd
    # envsubst < /etc/nginx/conf.d/default.conf.template > default.conf
    # sudo chown root:root default.conf
    # sudo mv default.conf /etc/nginx/conf.d/default.conf
    # sudo nginx -s reload
    # sudo systemctl restart nginx.service
    
    
    echo "Creating internal network..."
    docker network create internal_${BUILD_ENVIRONMENT}
    
    echo "Pulling..."
    docker-compose -f ~/locg/deploy/docker-compose.yaml --project-name locg_"$BUILD_ENVIRONMENT" pull
    
    echo "Restarting..."
    docker-compose -f ~/locg/deploy/docker-compose.yaml --project-name locg_"$BUILD_ENVIRONMENT" up -d
    
    echo "Done"
}