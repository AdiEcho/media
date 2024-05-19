# slug

I found this code:

https://github.com/davegonzalez/ott-boilerplate/blob/master/actions.js

where you can do this instead:

~~~
GET /collections/my-dinner-with-andre/items?site_id=59054 HTTP/1.1
Host: api.vhx.com
Authorization: Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6ImQ2YmZlZmMzNGIyNTdhYTE4Y2E2...
~~~

you cannot change ANYTHING about this request other than the host:

~~~
200 OK https://api.vhx.com/collections/my-dinner-with-andre/items?site_id=59054
200 OK https://api.vhx.tv/collections/my-dinner-with-andre/items?site_id=59054
404 Not Found https://api.vhx.com/collections/my-dinner-with-andre
404 Not Found https://api.vhx.com/collections/my-dinner-with-andre/items
404 Not Found https://api.vhx.com/collections/my-dinner-with-andre?site_id=59054
404 Not Found https://api.vhx.com/v2/sites/59054/collections/my-dinner-with-andre
404 Not Found https://api.vhx.com/v2/sites/59054/collections/my-dinner-with-andre/items
404 Not Found https://api.vhx.com/v2/sites/59054/collections/my-dinner-with-andre/items?site_id=59054
404 Not Found https://api.vhx.com/v2/sites/59054/collections/my-dinner-with-andre?site_id=59054
404 Not Found https://api.vhx.com/v2/sites/59054/videos/my-dinner-with-andre
404 Not Found https://api.vhx.com/v2/sites/59054/videos/my-dinner-with-andre/items
404 Not Found https://api.vhx.com/v2/sites/59054/videos/my-dinner-with-andre/items?site_id=59054
404 Not Found https://api.vhx.com/v2/sites/59054/videos/my-dinner-with-andre?site_id=59054
404 Not Found https://api.vhx.com/videos/455774/items?site_id=59054
404 Not Found https://api.vhx.com/videos/my-dinner-with-andre
404 Not Found https://api.vhx.com/videos/my-dinner-with-andre/items
404 Not Found https://api.vhx.com/videos/my-dinner-with-andre/items?site_id=59054
404 Not Found https://api.vhx.com/videos/my-dinner-with-andre?site_id=59054
404 Not Found https://api.vhx.tv/collections/my-dinner-with-andre
404 Not Found https://api.vhx.tv/collections/my-dinner-with-andre/items
404 Not Found https://api.vhx.tv/collections/my-dinner-with-andre?site_id=59054
404 Not Found https://api.vhx.tv/v2/sites/59054/collections/my-dinner-with-andre
404 Not Found https://api.vhx.tv/v2/sites/59054/collections/my-dinner-with-andre/items
404 Not Found https://api.vhx.tv/v2/sites/59054/collections/my-dinner-with-andre/items?site_id=59054
404 Not Found https://api.vhx.tv/v2/sites/59054/collections/my-dinner-with-andre?site_id=59054
404 Not Found https://api.vhx.tv/v2/sites/59054/videos/my-dinner-with-andre
404 Not Found https://api.vhx.tv/v2/sites/59054/videos/my-dinner-with-andre/items
404 Not Found https://api.vhx.tv/v2/sites/59054/videos/my-dinner-with-andre/items?site_id=59054
404 Not Found https://api.vhx.tv/v2/sites/59054/videos/my-dinner-with-andre?site_id=59054
404 Not Found https://api.vhx.tv/videos/my-dinner-with-andre
404 Not Found https://api.vhx.tv/videos/my-dinner-with-andre/items
404 Not Found https://api.vhx.tv/videos/my-dinner-with-andre/items?site_id=59054
404 Not Found https://api.vhx.tv/videos/my-dinner-with-andre?site_id=59054
~~~
