#/bin/bash

gcloud config configurations create emulator 2>/dev/null || true
gcloud config set auth/disable_credentials true
gcloud config set project test-project
gcloud config set -q api_endpoint_overrides/spanner http://spanner:9020/
gcloud spanner instances describe test-instance 2>/dev/null || gcloud spanner instances create test-instance --config=emulator-config --description=Emulator --nodes=1
gcloud spanner databases describe test-database --instance=test-instance 2>/dev/null || gcloud spanner databases create test-database --instance=test-instance