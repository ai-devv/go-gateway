# Назначение

Сервис предназначен для возможности перенаправления на **любой** `redirect_uri` при использовании **любого** провайдера
`OAuth 2.0`

Допустим, вы решаете реализовать приложение, которое для входа использует некий `OAuth 2.0` провайдер.
При этом:

* приложение может располагаться на нескольких поддоменах;
* `OAuth 2.0` провайдер **НЕ** поддерживает несколько `redirect_uri`.

Т.е. приложение хочет, чтобы `OAuth 2.0` провайдер мог перенаправлять на:

* `http://subdomain0.domain.loc/oauth/callback`;
* `http://subdomain1.domain.loc/oauth/callback`.

Для решения этой проблемы можно использовать данный сервис. Суть следующая:

* приложение выполняет подготовительный запрос с указанием ссылки, на которую он хочет перенаправить пользователя

`POST` `http://gateway.loc/oauth/authorize`

```json
{
  "authorizeUrl": "http://oauth-provider.loc/oauth/authorize?client_id=client_id&redirect_uri=http%3A%2F%2Fsub0.app.loc%2Foauth%2Fcallback&response_type=code&scope=scope0%20scope1%20scope2&code_challenge=aDbPE7rEAOkQUHHNavRwhN-srU5eMCyUv-0k4BOvtz4&code_challenge_method=S256&state=app-state"
}
```

Получая в ответ данные, необходимые для перенаправления и обмена `code` на доступы в будущем:
```json
{
  "authorizeUrl": "http://oauth-provider.loc/oauth/authorize?client_id=client_id&code_challenge=aDbPE7rEAOkQUHHNavRwhN-srU5eMCyUv-0k4BOvtz4&code_challenge_method=S256&redirect_uri=http%3A%2F%2Fgateway.loc%2Foauth%2Fcallback&response_type=code&scope=scope0+scope1+scope2&state=gateway-1af43e2b-513a-49be-8988-c5f456465b2b",
  "redirectUrl": "http://gateway.loc/oauth/callback"
}
```

Сервис подменяет параметры `redirect_uri` и `state` на новые
* после возвращения из провайдера `OAuth 2.0` обратно в сервис (по `redirect_uri`) он перенаправляет пользователя
на изначальный `redirect_uri`, который был указан в первом запросе и заменяет параметр `state` на исходный

`GET` `http://gateway.loc/oauth/callback?code=external-code&state=gateway-1af43e2b-513a-49be-8988-c5f456465b2b`

Ответ:
`<a href="http://sub0.app.loc/oauth/callback?code=external-code&amp;state=app-state">Found</a>.`

Таким образом, во внешнем провайдере `OAuth 2.0` достаточно указать **одну** ссылку для перенаправления
(`redirect_uri`), при этом основная задача с перенаправлением на разные поддомены может быть решена

# TODO

- [X] Base logic
    - [X] OAuth authorize
    - [X] OAuth callback
- [ ] Readiness
- [ ] Closer
- [ ] Config
    - [ ] Support env files (???)
    - [ ] Live config/Static config (???)
- [ ] URL generating behind proxy
- [ ] Logging
- [ ] Redis
- [ ] Validate requests
- [ ] Project structure
    - [ ] Services
- [ ] Tests
- [ ] Open API (???)