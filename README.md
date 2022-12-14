# network_simulatiton
## ð About this repository
Tool for simulating network using MM1K and MD1K queue.

## ð How to run
### install
```shell
git clone git@github.com:yagikota/network_simulatiton.git
cd network_simulatiton
```
### simulate
```shell
go run main.go simulate -h
This CLI simulates queue.ðķ

Usage:
  network_simulation simulate [flags]

Flags:
  -k, --K int              capacity of service(capacity of queue and server) (default 50)
  -e, --end_time float     the end time of the simulation (default 100000)
  -h, --help               help for simulate
  -l, --lambda float       average arrival rate of a packet (default 0.5)
  -m, --myu float          average service rate of the server (default 1)
  -q, --queue_type int     the type of queue(MM1K:0, MD1K: 1)
  -s, --start_time float   the start time of the simulation
```

## ð References
* http://www.ieice-hbkb.org/files/05/05gun_01hen_06.pdf
* https://www.cc.nagasaki-u.ac.jp/sec_online_manual/fcpp/stdlib/stdug/general/11_3.htm
* [ãąãģããžãŦãŪčĻå·](https://ja.wikipedia.org/wiki/%E3%82%B1%E3%83%B3%E3%83%89%E3%83%BC%E3%83%AB%E3%81%AE%E8%A8%98%E5%8F%B7)
* https://qiita.com/SaitoTsutomu/items/f67c7e9f98dd27d94608
* https://ie.u-ryukyu.ac.jp/~asharif/pukiwiki/attach/%E3%82%B7%E3%83%9F%E3%83%A5%E3%83%AC%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3_mm1.pdf
* http://www.it-shikaku.jp/top30.php?hidari=01-02-06.php&migi=km01-02.php
* http://www.bunkyo.ac.jp/~nemoto/lecture/seisan/2001/queue1.pdf
* https://www.cis.nagasaki-u.ac.jp/labs/oguri/CompSimExm.pdf
* https://www.google.com/search?q=mm1k+%E5%BE%85%E3%81%A1%E8%A1%8C%E5%88%97+%E7%90%86%E8%AB%96%E5%80%A4&rlz=1C5CHFA_enJP1015JP1015&sxsrf=ALiCzsYUSgnIPu8j3TEUHsLM-cMcaxGyCA%3A1666036872497&ei=iLRNY8v5HYmnoATz0onwCA&ved=0ahUKEwjL3d7_huj6AhWJE4gKHXNpAo4Q4dUDCA8&uact=5&oq=mm1k+%E5%BE%85%E3%81%A1%E8%A1%8C%E5%88%97+%E7%90%86%E8%AB%96%E5%80%A4&gs_lcp=Cgdnd3Mtd2l6EAMyBwgAEB4QogQyBQgAEKIEMgUIABCiBDIFCAAQogQyBQgAEKIEOgQIIxAnOgUIIRCgAToECCEQFUoECE0YAUoECEEYAEoECEYYAFAAWK4jYMQnaABwAXgAgAGyAYgB-wqSAQMxLjmYAQCgAQHAAQE&sclient=gws-wiz
* http://www-optima.amp.i.kyoto-u.ac.jp/~takine/tmp/shiryou.pdf
* CLI
  * https://zenn.dev/yotto428/scraps/85385949f1304b
  * https://zenn.dev/tama8021/articles/22_0627_go_cobra_cli#go%E3%81%A7cli%E3%81%A3%E3%81%A6%E3%81%A9%E3%82%93%E3%81%AA%E3%81%97%E3%81%A6%E4%BD%9C%E3%82%8B%E3%81%AE%EF%BC%9F%EF%BC%9F
  * https://qiita.com/minamijoyo/items/cfd22e9e6d3581c5d81f
  * https://text.baldanders.info/golang/using-and-testing-cobra/
* Graph
  * https://github.com/gonum/plot
