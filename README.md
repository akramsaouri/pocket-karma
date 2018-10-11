# pocket-karma

## how to use pocket api 
1. Create new app and get your `consumer_key` from [Pocket: Developers](https://getpocket.com/developer/apps/) 
2. Get request token
```
curl -X POST \
  https://getpocket.com/v3/oauth/request \
  -H 'Content-Type: application/json' \
  -H 'X-Accept: application/json' \
  -d '{
	"consumer_key": "xxxx-yyyyyyyyyyyyyyyyyyyyyyyy",
	"redirect_uri": "app-name:authorizationFinished"
}'
```

2. Visit link to authorise app
`https://getpocket.com/auth/authorize?request_token=xxxxxxx-xxxx-xxxx-xxxx-xxxxxx&redirect_uri=karma-counter:authorizationFinished`

3. Exchange request token with access token
```
curl -X POST \
  https://getpocket.com/v3/oauth/authorize \
  -H 'Content-Type: application/json' \
  -H 'X-Accept: application/json' \
  -d '{
	"consumer_key":"xxxx-yyyyyyyyyyyyyyyyyyyyyyyy",
	"code":"xxxxxxx-xxxx-xxxx-xxxx-xxxxxx"
}'
```

## TODO 
[] continuous integration for up apex
