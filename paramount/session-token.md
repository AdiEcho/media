# session-token

~~~
sources\p495os\InterfaceC18648b.java
@InterfaceC16185f("/apps-api/{apiVersion}/{deviceType}/irdeto-control/session-token.json")
AbstractC18878l<DRMSessionEndpointResponse> m63002e(
   @InterfaceC16198s("apiVersion") String apiVersion,
   @InterfaceC16198s("deviceType") String deviceType,
   @InterfaceC16200u Map<String, String> drmSessionDetails,
   @InterfaceC16188i("Cache-Control") String cacheControl
);

@InterfaceC16185f("/apps-api/{apiVersion}/{deviceType}/irdeto-control/anonymous-session-token.json")
AbstractC18878l<DRMSessionEndpointResponse> m63020n(
   @InterfaceC16198s("apiVersion") String apiVersion,
   @InterfaceC16198s("deviceType") String deviceType,
   @InterfaceC16200u Map<String, String> drmSessionDetails,
   @InterfaceC16188i("Cache-Control") String cacheControl
);
~~~

call m63020n:

~~~
sources\com\viacbs\android\pplus\data\source\internal\domains\C14140h.java
public AbstractC18878l mo45715H0(HashMap drmSessionDetails, boolean z11) {
  int i11;
  AbstractC16545t.m52679i(drmSessionDetails, "drmSessionDetails");
  InterfaceC18648b interfaceC18648b = (InterfaceC18648b) this.f38427b.m48260b();
  String str = this.f38426a.get();
  String m48253d = this.f38428c.m48253d();
  InterfaceC14938c interfaceC14938c = this.f38429d;
  if (z11) {
      i11 = 0;
  } else {
      i11 = 10;
  }
  return interfaceC18648b.m63020n(str, m48253d, drmSessionDetails, interfaceC14938c.get(i11));
~~~

call mo45715H0:

~~~
sources\com\cbs\sc2\drm\DrmSessionManagerImpl.java
HashMap m52344m;
AbstractC16545t.m52679i(dataHolder, "dataHolder");
m52344m = AbstractC16467n0.m52344m(AbstractC15290i.m49171a("contentId", m15132p(dataHolder)));
if (this.f11283e.mo60934a()) {
   m52344m.put(AnalyticsAttribute.APPLICATION_PLATFORM_ATTRIBUTE, this.f11282d.mo44190n().m63523b());
   m52344m.put("version", this.f11281c.getAppVersionName());
   m52344m.put("model", this.f11281c.getDeviceName());
   m52344m.put("firmwareVersion", this.f11281c.getOsReleaseName());
}
if (m15140x()) {
   AbstractC18878l mo45715H0 = this.f11280b.mo45715H0(m52344m, z11);

sources\com\cbs\sc2\drm\DrmSessionManagerImpl$getDrmSessionWrapper$2.java
responseModel = (ResponseModel) interfaceC14939d2.mo45715H0(hashMap, this.$isRefreshLicense).m63718d();
~~~
