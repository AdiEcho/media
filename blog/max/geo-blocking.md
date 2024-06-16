# geo-blocking

for US web we get:

~~~
GET https://default.beam-any.prd.api.max.com/labs/api/v1/sessions/feature-flags/decisions?projectId=1172348a-ca10-4eaf-9a87-0dc246fd07b0 HTTP/2.0
cookie: st=eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiJ0b2tlbi03ZTgyMDViMS0wOGZmLTQyYzEtYj...
~~~

for US Android we get:

~~~
GET https://default.beam-any.prd.api.discomax.com/labs/api/v1/sessions/feature-flags/decisions?projectId=bc2a7695-6749-4cf8-b949-25b810098ab8&ignoreDisabled=true HTTP/2.0
authorization: Bearer eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiJ0b2tlbi0wYTAxODg2NS01ZjA...
~~~

and:

~~~
POST https://default.any-any.prd.api.discomax.com/labs/api/v1/sessions/feature-flags/decisions HTTP/2.0
authorization: Bearer eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiJ0b2tlbi0wYTAxODg2NS01ZjA...

{"projectId":"d8665e86-8706-415d-8d84-d55ceddccfb5"}
~~~

for UK web we get:

~~~
POST https://default.beam-any.prd.api.max.com/labs/api/v1/sessions/feature-flags/decisions HTTP/2.0
cookie: st=eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiJ0b2tlbi01OTEzMGU1Yy1mN2NiLTQ2NjItYT...

{"projectId":"67e7aa0f-b186-4b85-9cb0-86d40a23636c"}
~~~

and:

~~~
POST https://default.beam-any.prd.api.max.com/labs/api/v1/sessions/feature-flags/decisions HTTP/2.0
cookie: st=eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiJ0b2tlbi01OTEzMGU1Yy1mN2NiLTQ2NjItYT...

{
   "context":{"domain":"max.com"},
   "projectId":"4b93cd5e-e880-4831-9f3a-86c2f7a90691"
}
~~~

and:

~~~
GET https://default.beam-any.prd.api.max.com/labs/api/v1/sessions/feature-flags/decisions?projectId=1172348a-ca10-4eaf-9a87-0dc246fd07b0 HTTP/2.0
cookie: st=eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiJ0b2tlbi04ZThmODk0MC01ZTdjLTQ2YTAtOG...
~~~
