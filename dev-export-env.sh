for dir in mongodb gridfs rabbitmq postgres logger
do
    export $(grep -v '^#' $dir/.env | xargs -d '\n') 
done
