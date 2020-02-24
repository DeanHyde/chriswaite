#!/bin/bash

docker build . -t eu.gcr.io/cotswoldtrees/chris:latest

docker push eu.gcr.io/cotswoldtrees/chris:latest

gcloud run deploy chris --image eu.gcr.io/cotswoldtrees/chris:latest --project=cotswoldtrees --platform managed --region europe-west1