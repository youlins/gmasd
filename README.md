# what's gmasd
* gmasd: m2m access service realized by GO.
* apsn:id of gmasd
* sn:id of machine

# procotol : gmasd - data center
#### login
```
  action      gmasd - data center    msg
  connect           ->          connect [size]\r\npayload\r\n
  active            ->          active apsn sn [size]\r\npayload\r\n
  response          <-          [nok|ok] apsn sn [size]\r\npayload\r\n
  disactive         ->          disactive apsn sn keeptime[size]\r\nreason\r\n
```

#### report
```
  action      gmasd - data center    msg
  report            ->          report msg_name apsn sn [size]\r\nmsg_payload\r\n
```

#### config 
```
  action      gmasd - data center    msg
  support           ->              support config config_name apsn
  set               <-              set config_name apsn sn [size]\r\nmsg_payload\r\n
  get               <-              get config_name apsn sn [size]\r\n
  value             ->              value config_name apsn sn [size]\r\nconfig_result\r\n
```

#### query
```
  action      gmasd - data center    msg
  query             ->              query query_name apsn [sn] [size]\r\nmsg_payload\r\n
  result            <-              result query_name apsn [sn] [size]\r\nmsg_payload\r\n
```

# sdk
...
# example 
...
# summary
...
