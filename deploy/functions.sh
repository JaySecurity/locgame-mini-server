export ENVIRONMENT
export VERSION
export SERVICE_PORT
export REST_PORT
export REDIS_PORT
export NATS_PORT


function get_available_port() {
    /usr/bin/python3 -c 'import socket; s=socket.socket(); s.bind(("", 0)); print(s.getsockname()[1]); s.close()'
}


function locg_delete_environment() {
    echo "Deleting Existing Environment... "
    ENVIRONMENT=$1
    
    SERVICE_PORT=$(docker inspect --format "{{index (index (index .NetworkSettings.Ports \"8080/tcp\") 0) \"HostPort\"}}" locg_"${ENVIRONMENT}"_service_1)
    REST_PORT=$(docker inspect --format "{{index (index (index .NetworkSettings.Ports \"8080/tcp\") 0) \"HostPort\"}}" locg_"${ENVIRONMENT}"_rest-api_1)
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
    
    
    SERVICE_PORT=8500
    REST_PORT=9500
    REDIS_PORT=$(get_available_port)
    NATS_PORT=$(get_available_port)
    
    echo "Sleeping for 10 seconds…"
    sleep 10
    
    # shellcheck disable=SC2143
    if [[ $(docker ps -a | grep locg_"${BUILD_ENVIRONMENT}"_service_1) ]]; then
        echo "BUILD_Environment already exists. Deploying with saving the last ports."
        # SERVICE_PORT=$(docker inspect --format "{{index (index (index .NetworkSettings.Ports \"8080/tcp\") 0) \"HostPort\"}}" locg_"${BUILD_ENVIRONMENT}"_service_1)
        # REST_PORT=$(docker inspect --format "{{index (index (index .NetworkSettings.Ports \"8080/tcp\") 0) \"HostPort\"}}" locg_"${BUILD_ENVIRONMENT}"_rest-api_1)
        REDIS_PORT=$(docker inspect --format "{{index (index (index .NetworkSettings.Ports \"6379/tcp\") 0) \"HostPort\"}}" locg_"${BUILD_ENVIRONMENT}"_redis_1)
        NATS_PORT=$(docker inspect --format "{{index (index (index .NetworkSettings.Ports \"4222/tcp\") 0) \"HostPort\"}}" locg_"${BUILD_ENVIRONMENT}"_nats_1)
        
        if [[ "$SERVICE_PORT" =~ ^[0-9]+$ ]]; then
            echo "HTTP port: OK"
        else
            echo "Unable get HTTP port: $SERVICE_PORT"
            SERVICE_PORT=$(get_available_port)
        fi
        
        if [[ "$REST_PORT" =~ ^[0-9]+$ ]]; then
            echo "REST API port: OK"
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
    
    echo SERVICE: $SERVICE_PORT
    echo REST: $REST_PORT
    
    echo "Creating internal network..."
    docker network create internal_${BUILD_ENVIRONMENT}
    
    echo "Pulling..."
    docker-compose -f ~/locg/deploy/docker-compose.yaml --project-name locg_"$BUILD_ENVIRONMENT" pull
    
    echo "Restarting..."
    docker-compose -f ~/locg/deploy/docker-compose.yaml --project-name locg_"$BUILD_ENVIRONMENT" up -d
    
    echo "Done"
}