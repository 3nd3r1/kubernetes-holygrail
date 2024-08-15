#!/usr/bin/env bash
set -e

if [ -z $GCP_SA_NAME ]
then
  echo "GCP_SA_NAME is not set"
  exit 1
fi

if [ -z $GCP_SA_KEY_FILE ]
then
  echo "GCP_SA_KEY_FILE is not set"
  exit 1
fi

if [ -z $GCP_BUCKET_NAME ]
then
  echo "GCP_BUCKET_NAME is not set"
  exit 1
fi

if [ -z $POSTGRES_USER ]
then
  echo "POSTGRES_USER is not set"
  exit 1
fi

if [ -z $POSTGRES_PASSWORD ]
then
  echo "POSTGRES_PASSWORD is not set"
  exit 1
fi

if [ -z $POSTGRES_DATABASE ]
then
  echo "POSTGRES_DATABASE is not set"
  exit 1
fi

if [ -z $POSTGRES_HOST ]
then
  echo "POSTGRES_HOST is not set"
  exit 1
fi

if [ -z $POSTGRES_PORT ]
then
  echo "POSTGRES_PORT is not set"
  exit 1
fi

if [ -z $NAMESPACE ]
then
  echo "NAMESPACE is not set"
  exit 1
fi

echo "Starting backup agent"

echo "Authenticating GCP service account..."
gcloud auth activate-service-account $GCP_SA_NAME --key-file $GCP_SA_KEY_FILE
echo "Authenticated GCP service account"

echo "Creating backup..."
timestamp=$(date +%Y-%m-%d-%H-%M-%S)
export PGPASSWORD=$POSTGRES_PASSWORD
pg_dump -h $POSTGRES_HOST -p $POSTGRES_PORT -U $POSTGRES_USER -d $POSTGRES_DATABASE > /usr/src/app/backup-${NAMESPACE}-${timestamp}.sql
echo "Backup created"

echo "Uploading backup..."
gcloud storage cp /usr/src/app/backup-${NAMESPACE}-${timestamp}.sql gs://$GCP_BUCKET_NAME
echo "Backup uploaded"

echo "Backup agent done"
