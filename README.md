# Table of Contents

- [1.0 Game Summary](#10-game-summary)
  - [1.1 Platforms](#11-platforms)
  - [1.2 Basic User Experience](#12-basic-user-experience)
- [2.0 Development Environments](#20-development-environments)
  - [2.1 Client](#21-client)
  - [2.2 Server](#22-server)
  - [2.3 Code Dependencies](#23-code-dependencies)
  - [3.0 Client Architecture](#30-client-architecture)
  - [3.1 UI System](#31-ui-system)
  - [3.2 Audio Engine](#32-audio-engine)
  - [3.3 Networking](#33-networking)
- [4.0 Server Architecture](#40-server-architecture)
  - [4.1 Network protocol](#41-network-protocol)
  - [4.2 Generating network code](#42-generating-network-code)
  - [4.3 Database structure](#43-database-structure)
  - [4.4 Game configuration system](#44-game-configuration-system)
  - [4.5 Message broker](#45-message-broker)
  - [4.6 CronJob Service](#46-cronjob-service)
- [5.0 Device Technical Requirements](#50-device-technical-requirements)
- [6.0 Build Delivery](#60-build-delivery)
- [7.0 User Interface](#70-user-interface)
  - [7.1 Localization](#71-localization)
- [8.0 Versioning](#80-versioning)
- [9.0 Server Maintenance](#90-server-maintenance)
- [10.0 Performance Testing](#100-performance-testing)
- [11.0 Server Targets](#110-server-targets)
- [12.0 Software and Data Version Control](#120-software-and-data-version-control)
- [13.0 Admin Panel](#130-admin-panel)
  - [13.1 Overview](#131-overview)
  - [13.2 Configuration System](#132-configuration-system)
  - [13.3 UI](#133-ui)

# **1.0 Game Summary**

Legends of Crypto is a card game using themes from well-known persons and influencers from Crypto World. Game is targeting WebGL, iOS and Android platforms.

The game is server-authoritative, which means all user / game data is stored on a server.

Additional game content will be delivered ODR ( On Demand Resources ).

## **1.1 Platforms**

LoC is targeting WebGL and mobile devices on Android and iOS. Support for other platforms is not currently required.

## **1.2 Basic User Experience**

This experience will change during the milestone process as features are added, but the expected final user flow will be:

- When the user starts the application they see a splash screen with the Project logo.
  - At this point the client performs a version check with the server.
  - The custom Loading Page follows with the content load indicator.
  - It will then prompt user to authenticate via MetaMask.
  - Load user data and progress upon successful authorization.
  - After loading, the user sees the game’s main page, where they have the opportunity to start the game.
- From this point, user will be available to use any game features that are not currently blocked by design.
- The game flow consists of repeating cycle:
  - The player participates in battle levels and receives currencies.
  - The player buys new packs of cards or spends currency.
  - The player participates in arena and tournaments.
- Other non-gameplay systems and flows may include:
  - Offers
  - Purchase flow
  - NFTs collecting

# **2.0 Development Environments**

## **2.1 Client**

The client game engine will be created in Unity 2021.3, using C# as the primary coding language. Specific platform requirements / libraries will be in Objective C, C++ and Java.

### Additional Platform Development Requirements

**iOS**: Mac computer running minimum OS X 10.14.6 version and Xcode 11.6 or higher  
**Android**: Android SDK and Java Development Kit (JDK 1.8); IL2CPP scripting backend requires Android NDK (r21b).

The hardware being used to develop the game consists solely of Mac computers that meet the Unity 2021 requirements.

**Unity 2019.4 requires**:
- OS: Mac OS High Sierra 10.13+
- Xcode 9.4 or later
- Android SDK and Java Development Kit (JDK); IL2CPP scripting backend requires Android NDK
- CPU: SSE2 instruction set support
- GPU: Graphics card with DX10 (shader model 4.0) capabilities

## **2.2 Server**

The server side will be created using the Go language. 

The Go language was chosen based on previous development experience, as well as due to its good multithreading, code execution speed, fast compilation and no need for a virtual machine. In general, this language is excellent for server-side development and is used by many large companies specifically for server-side development.
Go is a simple yet very powerful language with great documentation and well thought out built-in tools.
The language was created by Google as an alternative to C ++ and Java for application developers in the context of the company's tasks.

During development, a server with a Linux operating system and Docker installed is required.
This server must be accessible from the Internet and connected to MongoDB, NATS and Redis.

**Docker** is a set of platform as a service (PaaS) products that use OS-level virtualization to deliver software in packages called containers.

**MongoDB** is a source-available cross-platform document-oriented database program. Classified as a NoSQL database program, MongoDB uses JSON-like documents with optional schemas. 

**NATS** is an open-source messaging system (sometimes called message-oriented middleware). The NATS server is written in the Go programming language. Client libraries to interface with the server are available for dozens of major programming languages. The core design principles of NATS are performance, scalability, and ease of use.

**Redis** is an open source (BSD licensed), in-memory data structure store, used as a database, cache, and message broker.

There are no any features specific to NATS, it is used as a simple message broker. There are no technical hurdles switching to another message broker if necessary.

*What’s stored on the Server Side*:
- All calculations
- All player data
- All Game Configurations

This allows us to get rid of client side calculations and avoid possible hacking.

## **2.3 Code Dependencies**

# External
- **DOTween PRO 1.2.420** - Fast, efficient, fully type-safe object-oriented animation engine optimized for C#. Used for procedural animation.
- **2D Toolkit 4.3** - Efficient and flexible 2D sprite, collider set-up, text, tilemap, and UI system seamlessly integrated into the Unity environment. Used for powerful, flexible 2D in Unity and best performance optimization for Spine animations.
- **UniRx 7.1.0** - Reactive Extensions for Unity. Used for LINQ to asynchronous, LINQ to multithreading, and LINQ to events and more.
- **Spine Runtime 3.8** - C# and Unity3D runtime for Spine.
- **WalletConnect** - An extension of WalletConnectSharp that brings WalletConnect to Unity. Used for interactions with MetaMask.

# Internal
- **UI System** - see [UI System](#31-ui-system)
- **Addressables Manager** - toolset for working with Addressables System
- **Sound Manager** - toolset for working with sounds and music

# **3.0 Client Architecture**

## **3.1 UI System**

The UI System that will be utilized for the project was made internally and has successfully been used on other internal projects. This system is extremely easy to implement and provides a large list of functionality options. Additionally, because of our familiarity with the UI System, it can easily be updated and improved with ease.

Features
- Custom managers for UI element loading and switching (pages, popups, message boxes, pool objects)
- Simple scene management system
- Simple page management system
- Simple popup management system
- Simple message box management system
- Intelligent scene management
- Multi-layered camera system
- Completed and ready-to-go UI elements such as buttons, lists, toggles, etc.
- Many examples and demos
- No third-party resources are needed, everything works straight "out of the box."

## **3.2 Audio Engine**
The audio engine used will be Unity Audio System. It plays and mixes sounds of diverse formats on many operating systems. Throughout the game there is both music playback and sound effect playback. Background music plays throughout gameplay.

**Short Sounds**
- Compression format - Vorbis
- Load type - Compressed in memory
- Quality - 100%
- Stereo or mono (determined by context)

**Music**
- Compression format - Vorbis
- Load type - streaming
- Quality - 100%
- Stereo

## **3.3 Networking**

A custom network protocol based on WebSocket will be used. This is described in more detail in section [4.1 "Network Protocol"](#41-network-protocol).
Сode is generated for the client that allows you to call RPC methods.
When connected to the server, the WebSocket is processed in a separate thread.
This stream is divided into packets consisting of service information, method identifier and binary data of this method. Binary data are Protobuf messages.
If the player has lost the connection with the server, then a disconnect event will be triggered. You can handle this event as you like. At the moment, a window appears notifying the player about the loss of connection from the server and the "Reconnect" button. This action can be automated if required.
To avoid users spamming the reconnect button, the [backoff system](https://en.wikipedia.org/wiki/Exponential_backoff) will be implemented. In a couple of words, the backoff system will increase timeout with each sequential server call.
At the moment, request limiting is implemented on the server side. You can configure this limit through the environment variable RATE_LIMIT. Default value - 500.
In the case when the server requires a reboot or shutdown (for example, in the case of a redeploy), the server sends all the connected client a notification about the planned shutdown and begins to wait until the clients disconnect from the server. At this point, the server is not accepting new clients. If clients do not disconnect from the server for a certain period, then they are disconnected forcibly.

# **4.0 Server Architecture**

The server side will be created using the Go language. In the future, it will be possible to launch several instances on a large number of servers. At the moment, this is possible, but there are problems with the sequence of answers from the server during matches if the players are in one match, but on different servers. The message broker NATS is used for inter-server communication. This allows data to be synchronized between server instances, as well as notifying servers of events related to clients that are connected to other instances.     

MongoDB database is used to store player and other data. MongoDB is a source-available cross-platform document-oriented database program. Classified as a NoSQL database program, MongoDB uses JSON-like documents with optional schemas.    

A custom network protocol will be used to communicate between the client and the server.    

When a client logs in, session data for this client is created on the server. Session data is stored in RAM while the client is connected to the service. After disconnecting the client from the service, the session data is automatically deleted.    

## **4.1 Network protocol**

The network protocol allows communication between a client and a server using methods and structures as in conventional programming languages. Since the connection to the server is permanent, this allows on the server side to store the client's session data while it is connected to the server, as well as raise events at any time on the client side and call certain code when the client is disconnected from the server.    

WebSocket protocol makes it easy to understand that the client is connected to the server and when it is disconnected. With this, you can store session data only for the duration of the client's connection to the server, which significantly reduces the number of queries to the database. Also, the WebSocket protocol allows you to send any message to the client at any time, without any manipulation on the part of the client. 
That is why it was decided to implement a proprietary protocol that would take the best from analogs and get rid of their shortcomings. The network protocol is built on top of the WebSocket protocol. It’s based around the idea of specifying the methods that can be called remotely with their parameters and return types. On the server side, the server implements this interface and runs a server to handle client calls. On the client side, the client has a stub that provides the same methods as the server.
Protocol buffers (protobuf) are used for marshaling, Google’s mature open source mechanism for serializing structured data.     

Protocol buffers allow you to describe the data structures used in network communication between the server and the client. These structures are used as arguments to invoked methods and also as results in responses.    

This network protocol uses the following network packet structure:
<table>
<thead>
  <tr>
    <th colspan="7">Network Packet</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td colspan="4">Header</td>
    <td colspan="3">Data</td>
  </tr>
  <tr>
    <td>2 bytes</td>
    <td>4 bytes</td>
    <td>4 bytes</td>
    <td>2 bytes</td>
    <td>(Payload Size) bytes</td>
    <td colspan="2">(Error Size) bytes</td>
  </tr>
  <tr>
    <td rowspan="2">Method ID</td>
    <td rowspan="2">Sequence ID</td>
    <td rowspan="2">Payload size</td>
    <td rowspan="2">Error size</td>
    <td rowspan="2">Payload data</td>
    <td>2 bytes</td>
    <td>(Error Size - 2) bytes</td>
  </tr>
  <tr>
    <td>Error Code</td>
    <td>Error Description</td>
  </tr>
</tbody>
</table>

Upon receipt of this packet is called a particular method, which has been described through the definition file and has been assigned Method ID. The data itself is deserialized from Payload Data using Protobuf. Also, this structure allows you to indicate an error, if any. In such cases, an exception will be raised on the client side and must be handled.
If a network packet cannot be sent at one time, it is divided into several packets, and each packet is assigned an incremental Sequence ID value, with which you can combine a set of packets into one complete packet after receiving these packets.
After connecting to the server, the client is assigned a "context" in which certain information can be stored, as well as identify which client invoked a particular method without having to re-authenticate.

For the security of network data transmission, a TLS 1.2 connection is used (traffic is encrypted from both sides).

## **4.2 Generating network code**

Methods called on the server and client must be described in the definition file, specifying protobuf structures as arguments and call results. Using this definition file and a written utility called "network-gen", code is generated to communicate between the server and the client. After the code is generated, the methods will be available from the client and server side just like normal methods.

The definition file contains information about possible errors, RPC methods and methods on the client side that can be called at any time from the server side.

An example of defining an error in a definition file:

```
errors {
    // UnexpectedError is an error that occurs when an unknown error occurs for unknown reasons.
    UnexpectedError

    // AccountNotFound error that occurs when the required account is not found.
    AccountNotFound
}
```

Exceptions will be generated  on the client side and can be handled.
If exception handling is required, then they must be passed as arguments to the generated method. If a specific exception has not been handled, then the default exception handling will be triggered.

For example:
```c#
public static async Task<PlayerResources> BuyMarketItem(string itemId)
{
    return await _router.BuyMarketItem(new BuyMarketItem {ItemID = itemId}, typeof(NotEnoughResourcesException));
}
```
After passing the exception as an argument to the method, it needs to be handled at a higher level.

To define RPC methods, you need to write the name of the method, the structure that it takes as an argument, and the structure that will store the result.

For example:
```
server {
    // Login allows to perform the user authentication step.
    Login (LoginRequest) (LoginResponse)
    ....
}
```

On the client and server side, code will be generated to process these methods.
In the current server-side example, the method will be added to the interface that needs to be implemented.

```go
type ServerRouter interface {
    // Login allows to perform the user authentication step.
    Login(*network.Client, *LoginRequest) (*LoginResponse, error)
    ....
}
```

On the client side, it can be used like this:

```c#
// Login allows to perform the user authentication step.
public static async Task<LoginResponse> Login(string id, IdType type)
{
    return await _router.Login(new LoginRequest {DeviceID = DeviceId, ID = id, Type = type});
}
```

Using this code generation on the server side, you need to implement interface methods, and on the client side, you just call methods with passing arguments and processing the results of calling these methods.

To create methods called on the client side at any time from the server side, you need to define the method in the "client" section:

```
client {
    // OnRatingChanged is called when the player's rating changes
    OnRatingChanged (GlobalLevelInfo)
    ...
}
```

These methods have no return result, they are used as events to notify the client about.

On the client side, an interface will be generated that needs to be implemented:

```c#
public interface IClientRouter
{
    // OnRatingChanged is called when the player's rating changes
    void OnRatingChanged(RatingInfo arg, Error error);
    ...
}
```

On the server side, sending an event to the client would look like this:

```go
s.clientHandler.OnRatingChanged(session.GetClient(), &arena.RatingInfo{
	Rating: newGlobalLevel
}, nil)
```

Generate code using this command, passing specific arguments and the path to the definition file:

```bash
network-gen -go_server_out=pkg/model/proto -csharp_client_out=${LOCG_PROTO_UNITY_PATH} proto/handlers.network
```

## **4.3 Database structure**

The database stores many collections, which can be divided into two categories: frequently used and not frequently used.
The data that is often used is stored in the session and is located in the collection: `players`. This collection mainly stores information about the account, the player's arena data, the player's resources, tutorial data and decks. This information is almost always required, therefore it is stored in one collection.

On the server side, the data in this collection is not fully updated, but partially - only what has changed is changed. For this, a method was written that converts any protobuf structure into a bson map with changes.

## **4.4 Game configuration system**

All configurations are stored in a separate Git repository. Each config is stored in a separate yaml file. A repository can contain multiple branches for different environments.    

The server and Admin Panel at startup try to find a specific branch in the repository for the current environment (the name of the branch must match the name of the environment), if there is no such branch, the default branch will be used. After finding the desired branch, the process of getting all the changes from the repository will take place.    

This system allows us to:
- Have a history of changes and information about the people who changed the configuration.
- Use different configurations for different environments.
- Change configurations via Admin Panel.
- Ability to get configuration changes without rebuilding the server.

We can change configurations both through Admin Panel and manually.


## **4.5 Message broker**

A message broker is used to exchange messages (events) between servers.    
Messages are serialized data via Protobuf.     
Each message can be processed using event handlers.    
There are two types of event handlers:
- Events designed for a specific user (player channel).
- Events for all server instances.
#### Player Channel

When a player connects to the server, an automatic subscription to the player's channel (player.PLAYER_ID.*) occurs. This channel will receive events addressed to this particular player.    

A message addressed to a specific user can be sent like this:    
`pubsub.SendToPlayer(playerID, message)`

A message will be sent to the “player.PLAYER_ID.MESSAGE_TYPE” channel, where `PLAYER_ID` is the player's account ID and `MESSAGE_TYPE` is the name of the Protobuf structure.    

Each type of message needs a handler.    

To create a player channel message handler, you need to create a structure that will implement the interface:    
```go
type PlayerHandler interface {
    Handle(client *network.Client, clientHandler *proto.ClientHandler, data MessageData)
    GetDataType() reflect.Type
}
```

The Handle method will be called when a message with the data type returned by the GetDataType method is received.    
Each message handler needs to be registered:    
`pubsub.RegisterPlayerHandler(HANDLER)`

An example of sending and processing the event of the beginning of a new match is outlined below.    
First, create a handler which will implement the `PlayerHandler` interface:    
```go
type BeginMatchHandler struct {
    Sessions *sessions.SessionStore
}

func (h *BeginMatchHandler) Handle(client *network.Client, clientHandler *dto.ClientHandler, data pubsub.MessageData) {
    clientHandler.OnBeginMatch(client, data.(*game.BeginMatch), nil)
}

func (h *BeginMatchHandler) GetDataType() reflect.Type {
    return reflect.TypeOf(game.BeginMatch{})
}
```

Then register it:   
```go
pubsub.RegisterPlayerHandler(&BeginMatchHandler{Sessions: s.sessions }))
```

After that, it will be possible to send messages with the data type - game.BeginMatch:    

```go
pubsub.SendToPlayer(accountID, data)
```

#### Basic Message Handlers

To create a message handler addressed to all server instances, you need to create a structure and implement an interface:    

```go 
type Handler interface {
    Handle(data MessageData)
    GetDataType() reflect.Type
}
```

After that, this handler needs to be registered:    

```go
pubsub.Subscribe(channelName, handler)
```

## **4.6 CronJob Service**

CronJob Service is a service that allows you to schedule certain events to run periodically. It is used not only for CronJob, but also to perform difficult tasks, the call of which is initiated by receiving certain messages through a messenger broker.
This service is located next to the main service in the same repository, and uses the shared code. This allows us to use the existing code to work with configurations and connections with databases and a message broker.  
Based on the package https://github.com/robfig/cron    

Each task must implement the interface:   

```go
type Job interface {
    Run() error
    Init(config *config.Config, store *store.Store, logger *JobLogger)
    GetConfig() *config.Config
    GetStore() *store.Store
    GetLogger() *JobLogger

    GetPrefix() string
    GetName() string
    GetRetries() uint32

    ApplyOptions(options ...JobOption)
}
```

This interface implements the `BaseJob` structure.    
To create a new task, you need to create a structure that will include `BaseJob`:   
```go  
type SimpleJob struct {
    events.BaseJob
}

func (j *SimpleJob) Run() error {
    j.GetLogger().Info("Simple job")
    return nil
}
```

To register this task, you need to register it with the task name, time, number of repetitions, in case of failure of the task:   
```go
events.Register("*/1 * * * *",  &SimpleJob{})
```
Each task has access to all configurations and storages (databases).    
Each task is saved to the database and stored in the `RecurringJobData` structure:    
```go
message RecurringJobData {
    ObjectID ID = 1;
    string Name = 2;
    string Schedule = 3;
    Timestamp LastExecution = 4;
    uint32 SuccessCount = 5;
    uint32 ErrorCount = 6;
    bool Disabled = 7;
    uint32 Retries = 8;
    JobStatus Status = 9;
    Timestamp NextExecution = 10;
} 
```

When a specific task is started, a new instance of the JobData task is created:   
```go
message JobData {
    ObjectID ID = 1;
    ObjectID ParentJobID = 2;
    string Name = 3;
    uint32 Attempt = 4;
    JobStatus Status = 5;
    Timestamp StartedAt = 6;
    Timestamp FinishedAt = 7;
    string Output = 8;
}
```
This instance stores basic information about the task and information about the execution of the current instance of the task: Number of attempts, status, start and end times, and task execution log. `ParrentJobID` is the ID of the task (`RecurringJobData`) that created this task, i.e. during task registration, a `RecurringJobData` entry (the parent task) is created and each execution of this recurring task creates a `JobData` structure. This data is stored in a database.

Since all the main data is stored in the database, this allows us to display information about tasks in the Admin Panel.

## **5.0 Device Technical Requirements**

**iOS minimal requirements**:
- CPU Hexa-core 2.39 GHz
- 3 GB LPDDR2 of memory
- Apple GPU (three-core graphics)
- OS version 13.6

*Notes*:
Since iOS exclusions are based on the OS version (not hardware), setting the minimum to 13.6 would include iPhone 6S (2GB ram) and anything newer.

**WebGL minimal requirements**:
- PC with support of Chrome/Safari/Opera/Firefox Browsers
- 4 GB DDR of memory

**Android minimal requirements**:
- CPU Hexa-core 2.39 GHz
- 3 GB DDR of memory
- GPU Mali, Adreno, PowerVR
- OS version 4.4 (API level 19) minimal

**Android target requirements**:
- CPU Octa-core 2.8 GHz Kryo 385
- 8 GB DDR of memory
- GPU Adreno 630
- OS version 7.0 Nougat

*Notes*:
- Full resolution images
- High animation frame rate
- Fast loading

*Network*:
- Standing internet connection (3G, LTE, Wi-Fi) 

## **6.0 Build Delivery**

WebGL builds will be hosted at Amazon S3 and will be available in browser via direct link.

[Microsoft AppCenter](https://appcenter.ms/) will be used for Mobile build deliveries.

The App Center Build service helps developers build Android, iOS, macOS, and UWP apps using a secure cloud infrastructure. The repo can be connected in the App Center and be used to start building the app in the cloud on every commit. With App Center's other services, certain development aspects can be automated further:
Automatically release builds to testers and public app stores with App Center Distribute.
Run automated UI tests on thousands of real device and OS configurations in the cloud with App Center Test.
While the team is using GitLab as a project repo, CI/CD is mostly set up "out-of-the-box".
To trigger the build pipeline, the developer needs to push changes on respective branches. As of now, there are master and develop branches. Also, the build pipeline can be triggered via web interface.
Additionally, development is using Fastlane – an open source platform aimed at simplifying Android and iOS deployment.
Note: for internal iOS deployment, development is using internal certificates/provisioning profiles.

With correct AppCenter setup, builds are successfully deployed for both platforms - iOS (image 1) and Android (image 2).

// TODO ADD IMAGES

*Image 1 - AppCenter iOS releases*

*Image 2 - Android AppCenter releases*

After the build is successfully deployed, the respective application can be downloaded on the user’s device via the AppCenter application, URL, or by scanning QR code. Additionally, invites will be automatically sent to the respective person’s email.

## **7.0 User Interface**

The UI is divided into 3 independent layers:
- Screen - full screen menu
- Popup - next level over screen
- MessageBox - next level over popups (tutorial, information monologues/dialogs

UI Development Flow:
UI Artists receive functional mock-ups from the Design Department → UI Artists create assets from these mock-ups → Upon completion, assets are handed off from the UI Artist to either a Graphic UI artist or Engineer for implementation. 

## **7.1 Localization**

Remote Localization System:
- Generated from CSV file, result - asset (ScriptableObject), that can be updated remotely with no need of updating build. 
- Achieved by using Addressables System. 

As of now, there doesn’t seem to be a need for localized images. Default images with localized text components should be enough for project needs.

## **8.0 Versioning**

Use semantic versioning so there is a standard pattern to releases (MAJOR.MINOR.PATCH).

Latest compatible with server Game Client version will be stored on the server, Game Client will be forced to update when force update trigger will occur after the version compatibility process.

Compatibility between versions will be tested using Stage environment.

Every update, players will see in-game dialog with the message, “Please, update to the latest version” (or something similar). With this system in place, the game architecture is less complicated.

## **9.0 Server Maintenance**
At the server side, there is a toggle with functionality to enter/exit maintenance mode when needed. At the client side, there is a message box with a description about server maintenance. A player will not be able to enter a game until maintenance is over. So, there is a need to automate most of the processes to shorten maintenance time.

## **10.0 Performance Testing**

When developing software, performance testing plays an important role in determining how an application behaves on devices with minimal system requirements as well as devices with target system requirements.
Applications developed with Unity are easily performance-tested by utilizing the built-in Profiler Window. This greatly simplifies the development process in terms of performance optimization.

See below for a list of testing that will be applied throughout the development of the application.
- **Load Testing** - Used to determine the behavior of the system under different types of load.
- **Stress Testing** - Used to determine the reliability of an application under extreme load.
- **Soak Testing** - Used to determine the ability of an application to maintain the expected load in continuous mode and detect possible memory leaks.

The QA Department has an extensive arsenal of target devices. They are equipped to capture and report this data without issues. 

## **11.0 Server Targets**

- Response time < 100 ms
- No performance spikes

Estimated Drive Usage and RAM memory usage: 100 MB of available space, 350 MB RAM.
Platform targets have been determined based on our previous development experience and comp titles of the same genre. 

## **12.0 Software and Data Version Control**

Primary source of version control is in GIT.

Versioning game data will be implemented using [Git Tagging](https://git-scm.com/book/en/v2/Git-Basics-Tagging). Every launch, every delivery is tracked with an appropriate Tag.

For game data changes, every new live build will be tested at the Stage environment first for compatibility with the Production environment.

## **13.0 Admin Panel**

##  **13.1 Overview**
The Admin Panel uses the Go programming language. Frontend also uses the Go language with the [Vugu](https://github.com/vugu/vugu) library. Frontend code compiles to WebAssembly.    
The Admin panel requires a connection to the MongoDB database and a connection to a message broker through which communication with the main game services takes place.    

Connecting to the database is necessary for:
- editing player data
- editing the guilds and chat
- setting up messages for news
- storing data about administrators.    

Connection to the message broker is required to:
- notify servers about events (for example: technical break, configuration update), 
- notify players (user data changes and etc).    

The server receives configurations directly from the Git repository with the appropriate branch for the current environment.    

There is no state stored for the user, only his authentication. Each instance of the Admin Panel clones the repository to disk. Each time you change the configuration before pushing, new commits are received, in the future it is planned to add the ability to merge commits. If you run multiple instances of the Admin Panel, you may encounter situations with outdated information, this problem can be solved by periodically checking for new commits and getting them.     
At the moment, it is possible to manually update the repository to the latest state in the settings (for example, for situations where the state of the repository has been changed through the git client). If several users change data at the same time, then problems may arise. These problems can be solved by merging commits, but so far there is no such functionality.     
The Admin Panel does not contain local repository changes, all changes are immediately sent to the remote repository.    

Any disputes can be resolved using the standard git client.    

To work with Git, the [go-git](https://github.com/go-git/go-git) library is used, but unfortunately it does not currently support merging commits (only fast-forward), this functionality is planned by developers.    
If this functionality does not appear in the near future, then we will have to switch to other solutions (for example: [libgit2](https://libgit2.org)).    

To transfer configurations between environments, it is easier and faster to use the git client.    

### **13.2 Configuration System**
A configuration management system has been created. The configuration is stored in a separate repository.    
The admin panel at startup receives an external repository with configurations from the branch of the current environment.     
Each configuration change makes a commit to the repository. Through the admin panel, we can view the history of changes to certain features.      

Every restart of the admin panel retrieves all the changes in the repository, but we can also receive any new commits (made outside the admin panel) manually  in the admin panel settings.     

Updating a record will serialize the changes to yaml, write the file to disk (overwriting the previous record), and will then add, stage, and commit the changes to the remote repository.    

### **13.3 UI**
For the UI, the Vugu library is used. This allows developers to create UI in the Go programming language.    
To create a UI, developers can create *.vugu files in which they can write in HTML/Go/CSS.     
In HTML code, we can access data in Go, subscribe to DOM events, and call Go methods. We can call JS code from Go.      

Basic components have been created that allow us to quickly build a user interface.     
These include: Page, Modal, Tabs (vertical and horizontal), TextField, IntField, FloatField, BoolField, Select, TextArea, Tooltip, Card, Form, Buttons, Spacer, Image, Link, Span, Container, Table, Row, ButtonGroup, SearchField, Sidebar and others.    
The resulting code is compiled into WebAssembly and runs in-browser.    

