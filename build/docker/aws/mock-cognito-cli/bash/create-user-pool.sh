#!/bin/bash

# プロパティファイルを削除
if [ -f "/output/mock-cognito.properties" ]; then
    rm -f /output/mock-cognito.properties
fi

# ユーザープールの作成
echo execute create-user-pool
USER_POOL_ID=$(
  aws cognito-idp create-user-pool \
    --pool-name MyUserPool \
    --alias-attributes "email" \
    --username-attributes "email" \
    --query UserPool.Id \
    --output text \
    --endpoint-url ${ENDPOINT_URL} \
    --schema \
        Name=email,Required=true
)
echo USER_POOL_ID: $USER_POOL_ID

# アプリクライアントの作成
echo execute create-user-pool-client
set $(
  aws cognito-idp create-user-pool-client \
    --client-name MyUserPoolClient \
    --user-pool-id ${USER_POOL_ID} \
    --generate-secret \
    --query "[UserPoolClient.ClientId, UserPoolClient.ExplicitAuthFlows]" \
    --output text \
    --endpoint-url ${ENDPOINT_URL} \
    --explicit-auth-flows "ADMIN_USER_PASSWORD_AUTH"
)
CLIENT_ID=${1}
EXPLICIT_AUTH_FLOWS=${2}
echo CLIENT_ID: $CLIENT_ID
echo EXPLICIT_AUTH_FLOWS: $EXPLICIT_AUTH_FLOWS

# 管理者ユーザーの作成
echo execute admin-create-user
ADMIN_USER=$(aws cognito-idp admin-create-user \
  --user-pool-id ${USER_POOL_ID} \
  --username ${COGNITO_USER_EMAIL} \
  --message-action SUPPRESS \
  --desired-delivery-mediums EMAIL \
  --endpoint-url ${ENDPOINT_URL} \
  --user-attributes \
    Name=email,Value=${COGNITO_USER_EMAIL} \
    Name=email_verified,Value=true)

# 結果を出力
echo $ADMIN_USER

# 管理者ユーザーのパスワード設定
echo execute admin-set-user-password
ADMIN_SET_USER_PASSWORD=$(aws cognito-idp admin-set-user-password \
  --user-pool-id ${USER_POOL_ID} \
  --username ${COGNITO_USER_EMAIL} \
  --password ${COGNITO_USER_PASSWORD} \
  --permanent \
  --endpoint-url ${ENDPOINT_URL})

# 結果を出力
echo $ADMIN_SET_USER_PASSWORD

# ユーザーリスト、および、各ユーザーステータスの確認
# 出力結果のユーザーステータスに"CONFIRMED"があればOK
echo execute list-users
USER_LIST=$(aws cognito-idp list-users \
  --user-pool-id ${USER_POOL_ID} \
  --endpoint-url ${ENDPOINT_URL})

# 結果を出力
echo $USER_LIST

# プロパティファイルにユーザープールIDとクライアントIDを設定
cat <<EOF > /output/mock-cognito.properties
AWS_COGNITO_USER_POOL_ID=$USER_POOL_ID
AWS_COGNITO_USER_CLIENT_ID=$CLIENT_ID
EOF
