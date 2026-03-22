# What is the Real Time API?

SynchroGuard relies on a data stream to exchange data between the different
components of the system.

The data collected in the field are stored in a data stream, and consumed by the
different services part of SynchroGuard Real Time processing layer (i.e., state
estimation, grid event detection, fault location, topology processing)
to perform their tasks.

The results of the the Real Time processing layer are stored in data stream as
well, and consumed by other servoices, and made available to other services
(e.g., the data storing and scada integration services).

The data available in the data stream includes:

- Sensor Measurements: the data collected from devices in the field (such as PMUs,
  IEDs, etc.), and the [SCADA](../scada) system.
- Estimated Measurements: the measurements computed by the state estimation
  service, for not monitored nodes.
- Grid Events: the power quality and grid load events detected by the grid event
  detection service.
- Faults: the faults detected (and eventually located) by the fault location
  service.
- Topology changes: the changes in the grid topology detected by the topology
  proceessing service.

Optionally, the data stream can be configured to be consumed by external systems
as well, to make the data available to third party applications.

The Real Time API defines the data format of the data stream, and the protocol to
exchange data between the different components of the system, and with external
systems.
