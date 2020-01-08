package dms

const QPlayDeviceDesc = `<?xml version="1.0" encoding="utf-8"?>
<root xmlns="urn:schemas-upnp-org:device-1-0" xmlns:qq="http://www.tencent.com">
   <specVersion>
      <major>1</major>
      <minor>0</minor>
   </specVersion>
   <device>
      <deviceType>urn:schemas-upnp-org:device:MediaRenderer:1</deviceType>
      <friendlyName>Tx Media Renderer</friendlyName>
      <manufacturer>Tencent</manufacturer>
      <manufacturerURL />
      <modelDescription>Media Renderer Device</modelDescription>
      <modelName>Tx Media Renderer</modelName>;
      <modelNumber>1.0</modelNumber>
      <UDN>uuid:46552A6A-D878-472e-9B67-816607BEB3CE</UDN>
      <qq:X_QPlay_SoftwareCapability>QPlay:2</qq:X_QPlay_SoftwareCapability>
      <serviceList>
         <service>
            <serviceType>urn:schemas-tencent-com:service:QPlay:2</serviceType>
            <serviceId>urn:tencent-com:serviceId:QPlay</serviceId>
            <SCPDURL>_urn-schemas-upnp-org-service-QPlay_scpd.xml</SCPDURL>
            <controlURL>_urn-schemas-upnp-org-service-QPlay_control</controlURL>
            <eventSubURL>_urn-schemas-upnp-org-service-QPlay_event</eventSubURL>
         </service>
         <service>
            <serviceType>urn:schemas-upnp-org:service:AVTransport:1</serviceType>
            <serviceId>urn:upnp-org:serviceId:AVTransport</serviceId>
            <SCPDURL>_urn-schemas-upnp-org-service-AVTransport_scpd.xml</SCPDURL>
            <controlURL>_urn-schemas-upnp-org-service-AVTransport_control</controlURL>
            <eventSubURL>_urn-schemas-upnp-org-service-AVTransport_event</eventSubURL>
         </service>
         <service>
            <serviceType>urn:schemas-upnp-org:service:ConnectionManager:1</serviceType>
            <serviceId>urn:upnp-org:serviceId:ConnectionManager</serviceId>
            <SCPDURL>_urn-schemas-upnp-org-service-ConnectionManager_scpd.xml</SCPDURL>
            <controlURL>_urn-schemas-upnp-org-service-ConnectionManager_control</controlURL>
            <eventSubURL>_urn-schemas-upnp-org-service-ConnectionManager_event</eventSubURL>
         </service>
         <service>
            <serviceType>urn:schemas-upnp-org:service:RenderingControl:1</serviceType>
            <serviceId>urn:upnp-org:serviceId:RenderingControl</serviceId>
            <SCPDURL>_urn-schemas-upnp-org-service-RenderingControl_scpd.xml</SCPDURL>
            <controlURL>_urn-schemas-upnp-org-service-RenderingControl_control</controlURL>
            <eventSubURL>_urn-schemas-upnp-org-service-RenderingControl_event</eventSubURL>
         </service>
      </serviceList>
   </device>
</root>`
