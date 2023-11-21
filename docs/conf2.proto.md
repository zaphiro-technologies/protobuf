# Package: conf.v1

<div class="comment"><span><!-- markdownlint-disable --></span><br/><span>Messages describing PMU C37.118 Configurations. </span><br/><span>See [C37.118](https://www.typhoon-hil.com/documentation/typhoon-hil-software-manual/References/c37_118_protocol.html)</span><br/><span>protocol.</span><br/><span></span><br/></div>

## Imports

| Import | Description |
|--------|-------------|



## Options

| Name       | Value     | Description |
|------------|-----------|-------------|
| go_package | ./conf/v1 |             |




### Conf2Frame Diagram

```mermaid
classDiagram
direction LR

%% Configuration frame 2 
%% Headers used in rabbitMQ:
%% * `id`: id of the `Conf2Frame`
%% * `type`: always `Conf2Frame`
%% * `producerId`: the id of the producer (e.g. a PMU) linked to the configuration frame.
%% * `timestampId`: related measurement timestamp (if any)
%% 

class Conf2Frame {
  + uint32 DATA_RATE
  + List~Config~ configs
  + Conf2Header header
}
Conf2Frame --> `Config`
Conf2Frame --> `Conf2Header`

```
### Conf2Header Diagram

```mermaid
classDiagram
direction LR

%% Configuration frame 2 header

class Conf2Header {
  + uint32 FRACSEC
  + uint32 FRAMESIZE
  + uint32 IDCODE
  + uint32 NUM_PMU
  + uint32 SOC
  + uint32 SYNC
  + uint32 TIME_BASE
}

```
### Config Diagram

```mermaid
classDiagram
direction LR

%% Single PMU configuration according to Configuration frame 2

class Config {
  + uint32 ANNMR
  + List~uint32~ ANUNIT
  + uint32 CFGCNT
  + string CHNAM
  + uint32 DGNMR
  + List~uint32~ DIGUNIT
  + uint32 FNOM
  + uint32 FORMAT
  + uint32 IDCODE
  + uint32 PHNMR
  + List~uint32~ PHUNIT
  + string STN
}

```

## Message: Conf2Frame
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: conf.v1.Conf2Frame</div>

<div class="comment"><span>Configuration frame 2 </span><br/><span>Headers used in rabbitMQ:</span><br/><span>* `id`: id of the `Conf2Frame`</span><br/><span>* `type`: always `Conf2Frame`</span><br/><span>* `producerId`: the id of the producer (e.g. a PMU) linked to the configuration frame.</span><br/><span>* `timestampId`: related measurement timestamp (if any)</span><br/><span></span><br/></div>

| Field     | Ordinal | Type        | Label    | Description                   |
|-----------|---------|-------------|----------|-------------------------------|
| DATA_RATE | 3       | uint32      |          | Rate of data transmission     |
| configs   | 2       | Config      | Repeated | Set of PMU configurations     |
| header    | 1       | Conf2Header |          | Configuration frame 2 header  |




## Message: Conf2Header
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: conf.v1.Conf2Header</div>

<div class="comment"><span>Configuration frame 2 header</span><br/></div>

| Field     | Ordinal | Type   | Label | Description                                          |
|-----------|---------|--------|-------|------------------------------------------------------|
| FRACSEC   | 5       | uint32 |       | Fraction of Second and Message Time Quality          |
| FRAMESIZE | 2       | uint32 |       | Number of bytes in the frame                         |
| IDCODE    | 3       | uint32 |       | Stream source ID number                              |
| NUM_PMU   | 7       | uint32 |       | The number of PMUs included in the data frame        |
| SOC       | 4       | uint32 |       | SOC time stamp                                       |
| SYNC      | 1       | uint32 |       | Sync byte followed by frame type and version number  |
| TIME_BASE | 6       | uint32 |       | Resolution of FRACSEC time stamp                     |




## Message: Config
<div style="font-size: 12px; margin-top: -10px;" class="fqn">FQN: conf.v1.Config</div>

<div class="comment"><span>Single PMU configuration according to Configuration frame 2</span><br/></div>

| Field   | Ordinal | Type   | Label    | Description                            |
|---------|---------|--------|----------|----------------------------------------|
| ANNMR   | 5       | uint32 |          | Number of analog values                |
| ANUNIT  | 9       | uint32 | Repeated | Conversion factor for analog channels  |
| CFGCNT  | 12      | uint32 |          | Configuration change count             |
| CHNAM   | 7       | string |          | Phasor and channel names               |
| DGNMR   | 6       | uint32 |          | Number of digital status words         |
| DIGUNIT | 10      | uint32 | Repeated | Mask words for digital status words    |
| FNOM    | 11      | uint32 |          | Nominal line frequency code and flags  |
| FORMAT  | 3       | uint32 |          | Data format within data frame          |
| IDCODE  | 2       | uint32 |          | Data source ID number                  |
| PHNMR   | 4       | uint32 |          | Number of phasors                      |
| PHUNIT  | 8       | uint32 | Repeated | Conversion factor for phasor channels  |
| STN     | 1       | string |          | Station name                           |






<!-- Created by: Proto Diagram Tool -->
<!-- https://github.com/GoogleCloudPlatform/proto-gen-md-diagrams -->
