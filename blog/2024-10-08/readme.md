# paramount

this seems to work with everything:

<https://link.theplatform.com/s/dJ5BDC/media/guid/2198311517/esJvFlqdrcS_kFHnpxSuYp449E7tTexD?formats=MPEG-DASH&assetTypes=DASH_CENC_HDR10|DASH_CENC_PRECON>

if we omit `assetTypes`, we get multiple keys in some cases:

https://link.theplatform.com/s/dJ5BDC/media/guid/2198311517/Oo75PgAbcmt9xqqn1AMoBAfo190Cfhqi?formats=MPEG-DASH

if we omit `DASH_CENC_HDR10`, we get multiple Period in some cases:

<https://link.theplatform.com/s/dJ5BDC/media/guid/2198311517/esJvFlqdrcS_kFHnpxSuYp449E7tTexD?formats=MPEG-DASH&assetTypes=DASH_CENC_PRECON>

if we omit `DASH_CENC_PRECON`, we get `NoAssetTypeFormatMatches` in some cases
(France):

<https://link.theplatform.com/s/dJ5BDC/media/guid/2198311517/Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ?formats=MPEG-DASH&assetTypes=DASH_CENC_HDR10>

if we use just `DASH_CENC`, we get `NoAssetTypeFormatMatches` in some cases
(France):

<https://link.theplatform.com/s/dJ5BDC/media/guid/2198311517/Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ?formats=MPEG-DASH&assetTypes=DASH_CENC>

if we use `DASH_CENC|DASH_CENC_PRECON`, we get multiple Period in some cases:

<https://link.theplatform.com/s/dJ5BDC/media/guid/2198311517/esJvFlqdrcS_kFHnpxSuYp449E7tTexD?formats=MPEG-DASH&assetTypes=DASH_CENC|DASH_CENC_PRECON>

if we use `DASH_CENC_PRECON|DASH_CENC`, we get multiple Period in some cases:

<https://link.theplatform.com/s/dJ5BDC/media/guid/2198311517/esJvFlqdrcS_kFHnpxSuYp449E7tTexD?formats=MPEG-DASH&assetTypes=DASH_CENC_PRECON|DASH_CENC>

if we use `DASH_CENC_HDR10|DASH_CENC`, we get `NoAssetTypeFormatMatches` in
some cases (France):

<https://link.theplatform.com/s/dJ5BDC/media/guid/2198311517/Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ?formats=MPEG-DASH&assetTypes=DASH_CENC_HDR10|DASH_CENC>

if we use `DASH_CENC|DASH_CENC_HDR10`, we get `NoAssetTypeFormatMatches` in
some cases (France):

<https://link.theplatform.com/s/dJ5BDC/media/guid/2198311517/Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ?formats=MPEG-DASH&assetTypes=DASH_CENC|DASH_CENC_HDR10>
