package dms

const QPlayConnectionManagerDesc = `<?xml version="1.0" encoding="utf-8" ?>
<scpd xmlns="urn:schemas-upnp-org:service-1-0">
  <specVersion>
    <major>1</major>
    <minor>0</minor>
  </specVersion>
  <serviceStateTable>
    <stateVariable sendEvents="yes">
      <name>SourceProtocolInfo</name>
      <dataType>string</dataType>
    </stateVariable>
    <stateVariable sendEvents="yes">
      <name>SinkProtocolInfo</name>
      <dataType>string</dataType>
    </stateVariable>
  </serviceStateTable>
  <actionList>
    <action>
      <name>GetProtocolInfo</name>
      <argumentList>
    <argument>
      <name>Source</name>
      <direction>out</direction>
      <relatedStateVariable>SourceProtocolInfo</relatedStateVariable>
    </argument>
    <argument>
      <name>Sink</name>
      <direction>out</direction>
      <relatedStateVariable>SinkProtocolInfo</relatedStateVariable>
    </argument>
      </argumentList>
    </action>
  </actionList>
</scpd>`
