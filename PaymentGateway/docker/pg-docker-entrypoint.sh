#!/bin/bash

echo "Starting payment-gateway service"
if [[ "${INITIALIZE_PG}" == "true" ]]; then
  if [[ "${PP_ENV}" != "prod" ]]; then
    export UseTestApi=true
  else
    export UseTestApi=false
  fi
  cat /opt/${BASE_NAME}/config.json.tmpl | envsubst > /opt/${BASE_NAME}/config.json
fi
mkdir -p /opt/${BASE_NAME}/logs
function mark {
  match=$1
  file=$2
  mark=1
  while read -r data; do
    echo $data
    if [[ $data == *"$match"* ]]; then
      if [[ "$mark" == "1" ]]; then
        echo "done" >> $file
        mark=0
      fi
    fi
  done
}
BODY="{
    \"seed\": \"$seed\"
}"

curl -X POST -v "https://api-payment.magicwave.io/api/backend/setupAccount" \
  -u torpluse-api-payment-ny4KQDf3:acbhKTwXnPJRpcAMMXPRNJc7TaWx5busFqFP7beQYZLayYvtmQdEeGvAdRJwTbusNpuCnj7hu2kTSqmzLVQP7Tn4zFep7N2pTXnp \
	--data "${BODY}"
if [ $# -eq 0 ]
then
    /opt/${BASE_NAME}/payment-gateway | mark "Server is ready!" ".pg_ready" &> /opt/${BASE_NAME}/logs/payment.log
else
    exec "$@" | mark "Server is ready!" ".pg_ready"
fi
