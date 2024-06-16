# geo-blocking

UK web:

~~~
POST https://default.beam-any.prd.api.max.com/labs/api/v1/sessions/feature-flags/decisions HTTP/2.0
cookie: st=eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiJ0b2tlbi01OTEzMGU1Yy1mN2NiLTQ2NjItYT...

{"projectId":"67e7aa0f-b186-4b85-9cb0-86d40a23636c"}
~~~

from JavaScript:

~~~js
{
   [R.MB.DEV]:"6606d33b-54ca-45b4-b9b6-39f1d8724cd4",
   [R.MB.INT]:"70446554-c357-491f-8e07-2bb0dd2e163a",
   [R.MB.STG]:"44b3145d-3fa5-42dd-9b51-f0f143cb4a3d",
   [R.MB.PRD]:"67e7aa0f-b186-4b85-9cb0-86d40a23636c"
}
~~~

and this:

~~~
POST https://default.beam-any.prd.api.max.com/labs/api/v1/sessions/feature-flags/decisions HTTP/2.0
cookie: st=eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiJ0b2tlbi01OTEzMGU1Yy1mN2NiLTQ2NjItYT...

{
   "context":{"domain":"max.com"},
   "projectId":"4b93cd5e-e880-4831-9f3a-86c2f7a90691"
}
~~~

from JavaScript:

~~~json
{
  "environments": {
    "dev": {
      "GLOBAL_DOMAIN": "api.yellow.max-tests.com",
      "LABS_PROJECT_ID": "04306406-a08d-4256-94a5-10634b8fb70c"
    },
    "int": {
      "GLOBAL_DOMAIN": "api.orange.max-tests.com",
      "LABS_PROJECT_ID": "166ae657-904d-4c3b-8b05-6d4299d44ffe"
    },
    "stg": {
      "GLOBAL_DOMAIN": "api.blue.max-tests.com",
      "LABS_PROJECT_ID": "8dd73d52-8db0-47c9-9354-0ef82dac4335"
    },
    "prd": {
      "GLOBAL_DOMAIN": "api.max.com",
      "LABS_PROJECT_ID": "4b93cd5e-e880-4831-9f3a-86c2f7a90691"
    }
  }
}
~~~

and:

~~~
GET https://default.beam-any.prd.api.max.com/labs/api/v1/sessions/feature-flags/decisions?projectId=1172348a-ca10-4eaf-9a87-0dc246fd07b0 HTTP/2.0
cookie: st=eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiJ0b2tlbi04ZThmODk0MC01ZTdjLTQ2YTAtOG...
~~~

from JavaScript:

~~~json
{
  "environments": {
    "dev": {
      "boltGlobalDomain": "api.yellow.max-tests.com",
      "labsProjectId": "76e8c89e-ad0f-45f7-b147-dbf90380a1e2"
    },
    "int": {
      "boltGlobalDomain": "api.orange.max-tests.com",
      "labsProjectId": "f01c92bb-db5e-434c-9840-c1014f857070"
    },
    "stg": {
      "boltGlobalDomain": "api.blue.max-tests.com",
      "labsProjectId": "5f1dae91-10e8-4fed-9774-4a0cbd2a71cd"
    },
    "prd": {
      "boltGlobalDomain": "api.max.com",
      "labsProjectId": "1172348a-ca10-4eaf-9a87-0dc246fd07b0"
    }
  }
}
~~~
