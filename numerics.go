package qairc

var Numerics map[string]string

func initMap() {
	Numerics = make(map[string]string)
	Numerics["001"] = "RPL_WELCOME"                //RFC2812
	Numerics["002"] = "RPL_YOURHOST"               //RFC2812
	Numerics["003"] = "RPL_CREATED"                //RFC2812
	Numerics["004"] = "RPL_MYINFO"                 //RFC2812
	Numerics["004"] = "RPL_MYINFO"                 //KineIRCd
	Numerics["005"] = "RPL_BOUNCE"                 //RFC2812
	Numerics["005"] = "RPL_ISUPPORT"               //Also
	Numerics["006"] = "RPL_MAP"                    //Unreal
	Numerics["007"] = "RPL_MAPEND"                 //Unreal
	Numerics["008"] = "RPL_SNOMASK"                //ircu
	Numerics["009"] = "RPL_STATMEMTOT"             //ircu
	Numerics["010"] = "RPL_BOUNCE"                 //<hostname>
	Numerics["010"] = "RPL_STATMEM"                //ircu
	Numerics["014"] = "RPL_YOURCOOKIE"             //Hybrid?
	Numerics["015"] = "RPL_MAP"                    //ircu
	Numerics["016"] = "RPL_MAPMORE"                //ircu
	Numerics["017"] = "RPL_MAPEND"                 //ircu
	Numerics["042"] = "RPL_YOURID"                 //IRCnet
	Numerics["043"] = "RPL_SAVENICK"               //IRCnet
	Numerics["050"] = "RPL_ATTEMPTINGJUNC"         //aircd
	Numerics["051"] = "RPL_ATTEMPTINGREROUTE"      //aircd
	Numerics["200"] = "RPL_TRACELINK"              //RFC1459
	Numerics["201"] = "RPL_TRACECONNECTING"        //RFC1459
	Numerics["202"] = "RPL_TRACEHANDSHAKE"         //RFC1459
	Numerics["203"] = "RPL_TRACEUNKNOWN"           //RFC1459
	Numerics["204"] = "RPL_TRACEOPERATOR"          //RFC1459
	Numerics["205"] = "RPL_TRACEUSER"              //RFC1459
	Numerics["206"] = "RPL_TRACESERVER"            //RFC1459
	Numerics["207"] = "RPL_TRACESERVICE"           //RFC2812
	Numerics["208"] = "RPL_TRACENEWTYPE"           //RFC1459
	Numerics["209"] = "RPL_TRACECLASS"             //RFC2812
	Numerics["210"] = "RPL_TRACERECONNECT"         //RFC2812
	Numerics["210"] = "RPL_STATS"                  //aircd
	Numerics["211"] = "RPL_STATSLINKINFO"          //RFC1459
	Numerics["212"] = "RPL_STATSCOMMANDS"          //RFC1459
	Numerics["213"] = "RPL_STATSCLINE"             //RFC1459
	Numerics["214"] = "RPL_STATSNLINE"             //RFC1459
	Numerics["215"] = "RPL_STATSILINE"             //RFC1459
	Numerics["216"] = "RPL_STATSKLINE"             //RFC1459
	Numerics["217"] = "RPL_STATSQLINE"             //RFC1459
	Numerics["217"] = "RPL_STATSPLINE"             //ircu
	Numerics["218"] = "RPL_STATSYLINE"             //RFC1459
	Numerics["219"] = "RPL_ENDOFSTATS"             //RFC1459
	Numerics["220"] = "RPL_STATSPLINE"             //Hybrid
	Numerics["220"] = "RPL_STATSBLINE"             //Bahamut,
	Numerics["221"] = "RPL_UMODEIS"                //RFC1459
	Numerics["222"] = "RPL_MODLIST"                //
	Numerics["222"] = "RPL_SQLINE_NICK"            //Unreal
	Numerics["222"] = "RPL_STATSBLINE"             //Bahamut
	Numerics["223"] = "RPL_STATSELINE"             //Bahamut
	Numerics["223"] = "RPL_STATSGLINE"             //Unreal
	Numerics["224"] = "RPL_STATSFLINE"             //Hybrid,
	Numerics["224"] = "RPL_STATSTLINE"             //Unreal
	Numerics["225"] = "RPL_STATSDLINE"             //Hybrid
	Numerics["225"] = "RPL_STATSZLINE"             //Bahamut
	Numerics["225"] = "RPL_STATSELINE"             //Unreal
	Numerics["226"] = "RPL_STATSCOUNT"             //Bahamut
	Numerics["226"] = "RPL_STATSNLINE"             //Unreal
	Numerics["227"] = "RPL_STATSGLINE"             //Bahamut
	Numerics["227"] = "RPL_STATSVLINE"             //Unreal
	Numerics["228"] = "RPL_STATSQLINE"             //ircu
	Numerics["231"] = "RPL_SERVICEINFO"            //RFC1459
	Numerics["232"] = "RPL_ENDOFSERVICES"          //RFC1459
	Numerics["232"] = "RPL_RULES"                  //Unreal
	Numerics["233"] = "RPL_SERVICE"                //RFC1459
	Numerics["234"] = "RPL_SERVLIST"               //RFC2812
	Numerics["235"] = "RPL_SERVLISTEND"            //RFC2812
	Numerics["236"] = "RPL_STATSVERBOSE"           //ircu
	Numerics["237"] = "RPL_STATSENGINE"            //ircu
	Numerics["238"] = "RPL_STATSFLINE"             //ircu
	Numerics["239"] = "RPL_STATSIAUTH"             //IRCnet
	Numerics["240"] = "RPL_STATSVLINE"             //RFC2812
	Numerics["240"] = "RPL_STATSXLINE"             //AustHex
	Numerics["241"] = "RPL_STATSLLINE"             //RFC1459
	Numerics["242"] = "RPL_STATSUPTIME"            //RFC1459
	Numerics["243"] = "RPL_STATSOLINE"             //RFC1459
	Numerics["244"] = "RPL_STATSHLINE"             //RFC1459
	Numerics["245"] = "RPL_STATSSLINE"             //Bahamut,
	Numerics["246"] = "RPL_STATSPING"              //RFC2812
	Numerics["246"] = "RPL_STATSTLINE"             //ircu
	Numerics["246"] = "RPL_STATSULINE"             //Hybrid
	Numerics["247"] = "RPL_STATSBLINE"             //RFC2812
	Numerics["247"] = "RPL_STATSXLINE"             //Hybrid,
	Numerics["247"] = "RPL_STATSGLINE"             //ircu
	Numerics["248"] = "RPL_STATSULINE"             //ircu
	Numerics["248"] = "RPL_STATSDEFINE"            //IRCnet
	Numerics["249"] = "RPL_STATSULINE"             //Extension
	Numerics["249"] = "RPL_STATSDEBUG"             //Hybrid
	Numerics["250"] = "RPL_STATSDLINE"             //RFC2812
	Numerics["250"] = "RPL_STATSCONN"              //ircu,
	Numerics["251"] = "RPL_LUSERCLIENT"            //RFC1459
	Numerics["252"] = "RPL_LUSEROP"                //RFC1459
	Numerics["253"] = "RPL_LUSERUNKNOWN"           //RFC1459
	Numerics["254"] = "RPL_LUSERCHANNELS"          //RFC1459
	Numerics["255"] = "RPL_LUSERME"                //RFC1459
	Numerics["256"] = "RPL_ADMINME"                //RFC1459
	Numerics["257"] = "RPL_ADMINLOC1"              //RFC1459
	Numerics["258"] = "RPL_ADMINLOC2"              //RFC1459
	Numerics["259"] = "RPL_ADMINEMAIL"             //RFC1459
	Numerics["261"] = "RPL_TRACELOG"               //RFC1459
	Numerics["262"] = "RPL_TRACEPING"              //Extension
	Numerics["262"] = "RPL_TRACEEND"               //RFC2812
	Numerics["263"] = "RPL_TRYAGAIN"               //RFC2812
	Numerics["265"] = "RPL_LOCALUSERS"             //aircd,
	Numerics["266"] = "RPL_GLOBALUSERS"            //aircd,
	Numerics["267"] = "RPL_START_NETSTAT"          //aircd
	Numerics["268"] = "RPL_NETSTAT"                //aircd
	Numerics["269"] = "RPL_END_NETSTAT"            //aircd
	Numerics["270"] = "RPL_PRIVS"                  //ircu
	Numerics["271"] = "RPL_SILELIST"               //ircu
	Numerics["272"] = "RPL_ENDOFSILELIST"          //ircu
	Numerics["273"] = "RPL_NOTIFY"                 //aircd
	Numerics["274"] = "RPL_ENDNOTIFY"              //aircd
	Numerics["274"] = "RPL_STATSDELTA"             //IRCnet
	Numerics["275"] = "RPL_STATSDLINE"             //ircu,
	Numerics["276"] = "RPL_VCHANEXIST"             //
	Numerics["277"] = "RPL_VCHANLIST"              //
	Numerics["278"] = "RPL_VCHANHELP"              //
	Numerics["280"] = "RPL_GLIST"                  //ircu
	Numerics["281"] = "RPL_ENDOFGLIST"             //ircu
	Numerics["281"] = "RPL_ACCEPTLIST"             //
	Numerics["282"] = "RPL_ENDOFACCEPT"            //
	Numerics["282"] = "RPL_JUPELIST"               //ircu
	Numerics["283"] = "RPL_ALIST"                  //
	Numerics["283"] = "RPL_ENDOFJUPELIST"          //ircu
	Numerics["284"] = "RPL_ENDOFALIST"             //
	Numerics["284"] = "RPL_FEATURE"                //ircu
	Numerics["285"] = "RPL_GLIST_HASH"             //
	Numerics["285"] = "RPL_CHANINFO_HANDLE"        //aircd
	Numerics["285"] = "RPL_NEWHOSTIS"              //QuakeNet
	Numerics["286"] = "RPL_CHANINFO_USERS"         //aircd
	Numerics["286"] = "RPL_CHKHEAD"                //QuakeNet
	Numerics["287"] = "RPL_CHANINFO_CHOPS"         //aircd
	Numerics["287"] = "RPL_CHANUSER"               //QuakeNet
	Numerics["288"] = "RPL_CHANINFO_VOICES"        //aircd
	Numerics["288"] = "RPL_PATCHHEAD"              //QuakeNet
	Numerics["289"] = "RPL_CHANINFO_AWAY"          //aircd
	Numerics["289"] = "RPL_PATCHCON"               //QuakeNet
	Numerics["290"] = "RPL_CHANINFO_OPERS"         //aircd
	Numerics["290"] = "RPL_HELPHDR"                //Unreal
	Numerics["290"] = "RPL_DATASTR"                //QuakeNet
	Numerics["291"] = "RPL_CHANINFO_BANNED"        //aircd
	Numerics["291"] = "RPL_HELPOP"                 //Unreal
	Numerics["291"] = "RPL_ENDOFCHECK"             //QuakeNet
	Numerics["292"] = "RPL_CHANINFO_BANS"          //aircd
	Numerics["292"] = "RPL_HELPTLR"                //Unreal
	Numerics["293"] = "RPL_CHANINFO_INVITE"        //aircd
	Numerics["293"] = "RPL_HELPHLP"                //Unreal
	Numerics["294"] = "RPL_CHANINFO_INVITES"       //aircd
	Numerics["294"] = "RPL_HELPFWD"                //Unreal
	Numerics["295"] = "RPL_CHANINFO_KICK"          //aircd
	Numerics["295"] = "RPL_HELPIGN"                //Unreal
	Numerics["296"] = "RPL_CHANINFO_KICKS"         //aircd
	Numerics["299"] = "RPL_END_CHANINFO"           //aircd
	Numerics["300"] = "RPL_NONE"                   //RFC1459
	Numerics["301"] = "RPL_AWAY"                   //RFC1459
	Numerics["301"] = "RPL_AWAY"                   //KineIRCd
	Numerics["302"] = "RPL_USERHOST"               //RFC1459
	Numerics["303"] = "RPL_ISON"                   //RFC1459
	Numerics["304"] = "RPL_TEXT"                   //
	Numerics["305"] = "RPL_UNAWAY"                 //RFC1459
	Numerics["306"] = "RPL_NOWAWAY"                //RFC1459
	Numerics["307"] = "RPL_USERIP"                 //
	Numerics["307"] = "RPL_WHOISREGNICK"           //Bahamut,
	Numerics["307"] = "RPL_SUSERHOST"              //AustHex
	Numerics["308"] = "RPL_NOTIFYACTION"           //aircd
	Numerics["308"] = "RPL_WHOISADMIN"             //Bahamut
	Numerics["308"] = "RPL_RULESSTART"             //Unreal
	Numerics["309"] = "RPL_NICKTRACE"              //aircd
	Numerics["309"] = "RPL_WHOISSADMIN"            //Bahamut
	Numerics["309"] = "RPL_ENDOFRULES"             //Unreal
	Numerics["309"] = "RPL_WHOISHELPER"            //AustHex
	Numerics["310"] = "RPL_WHOISSVCMSG"            //Bahamut
	Numerics["310"] = "RPL_WHOISHELPOP"            //Unreal
	Numerics["310"] = "RPL_WHOISSERVICE"           //AustHex
	Numerics["311"] = "RPL_WHOISUSER"              //RFC1459
	Numerics["312"] = "RPL_WHOISSERVER"            //RFC1459
	Numerics["313"] = "RPL_WHOISOPERATOR"          //RFC1459
	Numerics["314"] = "RPL_WHOWASUSER"             //RFC1459
	Numerics["315"] = "RPL_ENDOFWHO"               //RFC1459
	Numerics["316"] = "RPL_WHOISCHANOP"            //RFC1459
	Numerics["317"] = "RPL_WHOISIDLE"              //RFC1459
	Numerics["318"] = "RPL_ENDOFWHOIS"             //RFC1459
	Numerics["319"] = "RPL_WHOISCHANNELS"          //RFC1459
	Numerics["320"] = "RPL_WHOISVIRT"              //AustHex
	Numerics["320"] = "RPL_WHOIS_HIDDEN"           //Anothernet
	Numerics["320"] = "RPL_WHOISSPECIAL"           //Unreal
	Numerics["321"] = "RPL_LISTSTART"              //RFC1459
	Numerics["322"] = "RPL_LIST"                   //RFC1459
	Numerics["323"] = "RPL_LISTEND"                //RFC1459
	Numerics["324"] = "RPL_CHANNELMODEIS"          //RFC1459
	Numerics["325"] = "RPL_UNIQOPIS"               //RFC2812
	Numerics["325"] = "RPL_CHANNELPASSIS"          //
	Numerics["326"] = "RPL_NOCHANPASS"             //
	Numerics["327"] = "RPL_CHPASSUNKNOWN"          //
	Numerics["328"] = "RPL_CHANNEL_URL"            //Bahamut,
	Numerics["329"] = "RPL_CREATIONTIME"           //Bahamut
	Numerics["330"] = "RPL_WHOWAS_TIME"            //
	Numerics["330"] = "RPL_WHOISACCOUNT"           //ircu
	Numerics["331"] = "RPL_NOTOPIC"                //RFC1459
	Numerics["332"] = "RPL_TOPIC"                  //RFC1459
	Numerics["333"] = "RPL_TOPICWHOTIME"           //ircu
	Numerics["334"] = "RPL_LISTUSAGE"              //ircu
	Numerics["334"] = "RPL_COMMANDSYNTAX"          //Bahamut
	Numerics["334"] = "RPL_LISTSYNTAX"             //Unreal
	Numerics["335"] = "RPL_WHOISBOT"               //Unreal
	Numerics["338"] = "RPL_CHANPASSOK"             //
	Numerics["338"] = "RPL_WHOISACTUALLY"          //ircu,
	Numerics["339"] = "RPL_BADCHANPASS"            //
	Numerics["340"] = "RPL_USERIP"                 //ircu
	Numerics["341"] = "RPL_INVITING"               //RFC1459
	Numerics["342"] = "RPL_SUMMONING"              //RFC1459
	Numerics["345"] = "RPL_INVITED"                //GameSurge
	Numerics["346"] = "RPL_INVITELIST"             //RFC2812
	Numerics["347"] = "RPL_ENDOFINVITELIST"        //RFC2812
	Numerics["348"] = "RPL_EXCEPTLIST"             //RFC2812
	Numerics["349"] = "RPL_ENDOFEXCEPTLIST"        //RFC2812
	Numerics["351"] = "RPL_VERSION"                //RFC1459
	Numerics["352"] = "RPL_WHOREPLY"               //RFC1459
	Numerics["353"] = "RPL_NAMREPLY"               //RFC1459
	Numerics["354"] = "RPL_WHOSPCRPL"              //ircu
	Numerics["355"] = "RPL_NAMREPLY_"              //QuakeNet
	Numerics["357"] = "RPL_MAP"                    //AustHex
	Numerics["358"] = "RPL_MAPMORE"                //AustHex
	Numerics["359"] = "RPL_MAPEND"                 //AustHex
	Numerics["361"] = "RPL_KILLDONE"               //RFC1459
	Numerics["362"] = "RPL_CLOSING"                //RFC1459
	Numerics["363"] = "RPL_CLOSEEND"               //RFC1459
	Numerics["364"] = "RPL_LINKS"                  //RFC1459
	Numerics["365"] = "RPL_ENDOFLINKS"             //RFC1459
	Numerics["366"] = "RPL_ENDOFNAMES"             //RFC1459
	Numerics["367"] = "RPL_BANLIST"                //RFC1459
	Numerics["368"] = "RPL_ENDOFBANLIST"           //RFC1459
	Numerics["369"] = "RPL_ENDOFWHOWAS"            //RFC1459
	Numerics["371"] = "RPL_INFO"                   //RFC1459
	Numerics["372"] = "RPL_MOTD"                   //RFC1459
	Numerics["373"] = "RPL_INFOSTART"              //RFC1459
	Numerics["374"] = "RPL_ENDOFINFO"              //RFC1459
	Numerics["375"] = "RPL_MOTDSTART"              //RFC1459
	Numerics["376"] = "RPL_ENDOFMOTD"              //RFC1459
	Numerics["377"] = "RPL_KICKEXPIRED"            //aircd
	Numerics["377"] = "RPL_SPAM"                   //AustHex
	Numerics["378"] = "RPL_BANEXPIRED"             //aircd
	Numerics["378"] = "RPL_WHOISHOST"              //Unreal
	Numerics["378"] = "RPL_MOTD"                   //AustHex
	Numerics["379"] = "RPL_KICKLINKED"             //aircd
	Numerics["379"] = "RPL_WHOISMODES"             //Unreal
	Numerics["380"] = "RPL_BANLINKED"              //aircd
	Numerics["380"] = "RPL_YOURHELPER"             //AustHex
	Numerics["381"] = "RPL_YOUREOPER"              //RFC1459
	Numerics["382"] = "RPL_REHASHING"              //RFC1459
	Numerics["383"] = "RPL_YOURESERVICE"           //RFC2812
	Numerics["384"] = "RPL_MYPORTIS"               //RFC1459
	Numerics["385"] = "RPL_NOTOPERANYMORE"         //AustHex,
	Numerics["386"] = "RPL_QLIST"                  //Unreal
	Numerics["386"] = "RPL_IRCOPS"                 //Ultimate
	Numerics["387"] = "RPL_ENDOFQLIST"             //Unreal
	Numerics["387"] = "RPL_ENDOFIRCOPS"            //Ultimate
	Numerics["388"] = "RPL_ALIST"                  //Unreal
	Numerics["389"] = "RPL_ENDOFALIST"             //Unreal
	Numerics["391"] = "RPL_TIME"                   //RFC1459
	Numerics["391"] = "RPL_TIME"                   //ircu
	Numerics["391"] = "RPL_TIME"                   //bdq-ircd
	Numerics["391"] = "RPL_TIME"                   //<server>
	Numerics["392"] = "RPL_USERSSTART"             //RFC1459
	Numerics["393"] = "RPL_USERS"                  //RFC1459
	Numerics["394"] = "RPL_ENDOFUSERS"             //RFC1459
	Numerics["395"] = "RPL_NOUSERS"                //RFC1459
	Numerics["396"] = "RPL_HOSTHIDDEN"             //Undernet
	Numerics["400"] = "ERR_UNKNOWNERROR"           //<command>
	Numerics["401"] = "ERR_NOSUCHNICK"             //RFC1459
	Numerics["402"] = "ERR_NOSUCHSERVER"           //RFC1459
	Numerics["403"] = "ERR_NOSUCHCHANNEL"          //RFC1459
	Numerics["404"] = "ERR_CANNOTSENDTOCHAN"       //RFC1459
	Numerics["405"] = "ERR_TOOMANYCHANNELS"        //RFC1459
	Numerics["406"] = "ERR_WASNOSUCHNICK"          //RFC1459
	Numerics["407"] = "ERR_TOOMANYTARGETS"         //RFC1459
	Numerics["408"] = "ERR_NOSUCHSERVICE"          //RFC2812
	Numerics["408"] = "ERR_NOCOLORSONCHAN"         //Bahamut
	Numerics["409"] = "ERR_NOORIGIN"               //RFC1459
	Numerics["411"] = "ERR_NORECIPIENT"            //RFC1459
	Numerics["412"] = "ERR_NOTEXTTOSEND"           //RFC1459
	Numerics["413"] = "ERR_NOTOPLEVEL"             //RFC1459
	Numerics["414"] = "ERR_WILDTOPLEVEL"           //RFC1459
	Numerics["415"] = "ERR_BADMASK"                //RFC2812
	Numerics["416"] = "ERR_TOOMANYMATCHES"         //IRCnet
	Numerics["416"] = "ERR_QUERYTOOLONG"           //ircu
	Numerics["419"] = "ERR_LENGTHTRUNCATED"        //aircd
	Numerics["421"] = "ERR_UNKNOWNCOMMAND"         //RFC1459
	Numerics["422"] = "ERR_NOMOTD"                 //RFC1459
	Numerics["423"] = "ERR_NOADMININFO"            //RFC1459
	Numerics["424"] = "ERR_FILEERROR"              //RFC1459
	Numerics["425"] = "ERR_NOOPERMOTD"             //Unreal
	Numerics["429"] = "ERR_TOOMANYAWAY"            //Bahamut
	Numerics["430"] = "ERR_EVENTNICKCHANGE"        //AustHex
	Numerics["431"] = "ERR_NONICKNAMEGIVEN"        //RFC1459
	Numerics["432"] = "ERR_ERRONEUSNICKNAME"       //RFC1459
	Numerics["433"] = "ERR_NICKNAMEINUSE"          //RFC1459
	Numerics["434"] = "ERR_SERVICENAMEINUSE"       //AustHex?
	Numerics["434"] = "ERR_NORULES"                //Unreal,
	Numerics["435"] = "ERR_SERVICECONFUSED"        //Unreal
	Numerics["435"] = "ERR_BANONCHAN"              //Bahamut
	Numerics["436"] = "ERR_NICKCOLLISION"          //RFC1459
	Numerics["437"] = "ERR_UNAVAILRESOURCE"        //RFC2812
	Numerics["437"] = "ERR_BANNICKCHANGE"          //ircu
	Numerics["438"] = "ERR_NICKTOOFAST"            //ircu
	Numerics["438"] = "ERR_DEAD"                   //IRCnet
	Numerics["439"] = "ERR_TARGETTOOFAST"          //ircu
	Numerics["440"] = "ERR_SERVICESDOWN"           //Bahamut,
	Numerics["441"] = "ERR_USERNOTINCHANNEL"       //RFC1459
	Numerics["442"] = "ERR_NOTONCHANNEL"           //RFC1459
	Numerics["443"] = "ERR_USERONCHANNEL"          //RFC1459
	Numerics["444"] = "ERR_NOLOGIN"                //RFC1459
	Numerics["445"] = "ERR_SUMMONDISABLED"         //RFC1459
	Numerics["446"] = "ERR_USERSDISABLED"          //RFC1459
	Numerics["447"] = "ERR_NONICKCHANGE"           //Unreal
	Numerics["449"] = "ERR_NOTIMPLEMENTED"         //Undernet
	Numerics["451"] = "ERR_NOTREGISTERED"          //RFC1459
	Numerics["452"] = "ERR_IDCOLLISION"            //
	Numerics["453"] = "ERR_NICKLOST"               //
	Numerics["455"] = "ERR_HOSTILENAME"            //Unreal
	Numerics["456"] = "ERR_ACCEPTFULL"             //
	Numerics["457"] = "ERR_ACCEPTEXIST"            //
	Numerics["458"] = "ERR_ACCEPTNOT"              //
	Numerics["459"] = "ERR_NOHIDING"               //Unreal
	Numerics["460"] = "ERR_NOTFORHALFOPS"          //Unreal
	Numerics["461"] = "ERR_NEEDMOREPARAMS"         //RFC1459
	Numerics["462"] = "ERR_ALREADYREGISTERED"      //RFC1459
	Numerics["463"] = "ERR_NOPERMFORHOST"          //RFC1459
	Numerics["464"] = "ERR_PASSWDMISMATCH"         //RFC1459
	Numerics["465"] = "ERR_YOUREBANNEDCREEP"       //RFC1459
	Numerics["466"] = "ERR_YOUWILLBEBANNED"        //RFC1459
	Numerics["467"] = "ERR_KEYSET"                 //RFC1459
	Numerics["468"] = "ERR_INVALIDUSERNAME"        //ircu
	Numerics["468"] = "ERR_ONLYSERVERSCANCHANGE"   //Bahamut,
	Numerics["469"] = "ERR_LINKSET"                //Unreal
	Numerics["470"] = "ERR_LINKCHANNEL"            //Unreal
	Numerics["470"] = "ERR_KICKEDFROMCHAN"         //aircd
	Numerics["471"] = "ERR_CHANNELISFULL"          //RFC1459
	Numerics["472"] = "ERR_UNKNOWNMODE"            //RFC1459
	Numerics["473"] = "ERR_INVITEONLYCHAN"         //RFC1459
	Numerics["474"] = "ERR_BANNEDFROMCHAN"         //RFC1459
	Numerics["475"] = "ERR_BADCHANNELKEY"          //RFC1459
	Numerics["476"] = "ERR_BADCHANMASK"            //RFC2812
	Numerics["477"] = "ERR_NOCHANMODES"            //RFC2812
	Numerics["477"] = "ERR_NEEDREGGEDNICK"         //Bahamut,
	Numerics["478"] = "ERR_BANLISTFULL"            //RFC2812
	Numerics["479"] = "ERR_BADCHANNAME"            //Hybrid
	Numerics["479"] = "ERR_LINKFAIL"               //Unreal
	Numerics["480"] = "ERR_NOULINE"                //AustHex
	Numerics["480"] = "ERR_CANNOTKNOCK"            //Unreal
	Numerics["481"] = "ERR_NOPRIVILEGES"           //RFC1459
	Numerics["482"] = "ERR_CHANOPRIVSNEEDED"       //RFC1459
	Numerics["483"] = "ERR_CANTKILLSERVER"         //RFC1459
	Numerics["484"] = "ERR_RESTRICTED"             //RFC2812
	Numerics["484"] = "ERR_ISCHANSERVICE"          //Undernet
	Numerics["484"] = "ERR_DESYNC"                 //Bahamut,
	Numerics["484"] = "ERR_ATTACKDENY"             //Unreal
	Numerics["485"] = "ERR_UNIQOPRIVSNEEDED"       //RFC2812
	Numerics["485"] = "ERR_KILLDENY"               //Unreal
	Numerics["485"] = "ERR_CANTKICKADMIN"          //PTlink
	Numerics["485"] = "ERR_ISREALSERVICE"          //QuakeNet
	Numerics["486"] = "ERR_NONONREG"               //
	Numerics["486"] = "ERR_HTMDISABLED"            //Unreal
	Numerics["486"] = "ERR_ACCOUNTONLY"            //QuakeNet
	Numerics["487"] = "ERR_CHANTOORECENT"          //IRCnet
	Numerics["487"] = "ERR_MSGSERVICES"            //Bahamut
	Numerics["488"] = "ERR_TSLESSCHAN"             //IRCnet
	Numerics["489"] = "ERR_VOICENEEDED"            //Undernet
	Numerics["489"] = "ERR_SECUREONLYCHAN"         //Unreal
	Numerics["491"] = "ERR_NOOPERHOST"             //RFC1459
	Numerics["492"] = "ERR_NOSERVICEHOST"          //RFC1459
	Numerics["493"] = "ERR_NOFEATURE"              //ircu
	Numerics["494"] = "ERR_BADFEATURE"             //ircu
	Numerics["495"] = "ERR_BADLOGTYPE"             //ircu
	Numerics["496"] = "ERR_BADLOGSYS"              //ircu
	Numerics["497"] = "ERR_BADLOGVALUE"            //ircu
	Numerics["498"] = "ERR_ISOPERLCHAN"            //ircu
	Numerics["499"] = "ERR_CHANOWNPRIVNEEDED"      //Unreal
	Numerics["501"] = "ERR_UMODEUNKNOWNFLAG"       //RFC1459
	Numerics["502"] = "ERR_USERSDONTMATCH"         //RFC1459
	Numerics["503"] = "ERR_GHOSTEDCLIENT"          //Hybrid
	Numerics["503"] = "ERR_VWORLDWARN"             //AustHex
	Numerics["504"] = "ERR_USERNOTONSERV"          //
	Numerics["511"] = "ERR_SILELISTFULL"           //ircu
	Numerics["512"] = "ERR_TOOMANYWATCH"           //Bahamut
	Numerics["513"] = "ERR_BADPING"                //ircu
	Numerics["514"] = "ERR_INVALID_ERROR"          //ircu
	Numerics["514"] = "ERR_TOOMANYDCC"             //Bahamut
	Numerics["515"] = "ERR_BADEXPIRE"              //ircu
	Numerics["516"] = "ERR_DONTCHEAT"              //ircu
	Numerics["517"] = "ERR_DISABLED"               //ircu
	Numerics["518"] = "ERR_NOINVITE"               //Unreal
	Numerics["518"] = "ERR_LONGMASK"               //ircu
	Numerics["519"] = "ERR_ADMONLY"                //Unreal
	Numerics["519"] = "ERR_TOOMANYUSERS"           //ircu
	Numerics["520"] = "ERR_OPERONLY"               //Unreal
	Numerics["520"] = "ERR_MASKTOOWIDE"            //ircu
	Numerics["520"] = "ERR_WHOTRUNC"               //AustHex
	Numerics["521"] = "ERR_LISTSYNTAX"             //Bahamut
	Numerics["522"] = "ERR_WHOSYNTAX"              //Bahamut
	Numerics["523"] = "ERR_WHOLIMEXCEED"           //Bahamut
	Numerics["524"] = "ERR_QUARANTINED"            //ircu
	Numerics["524"] = "ERR_OPERSPVERIFY"           //Unreal
	Numerics["525"] = "ERR_REMOTEPFX"              //CAPAB
	Numerics["526"] = "ERR_PFXUNROUTABLE"          //CAPAB
	Numerics["550"] = "ERR_BADHOSTMASK"            //QuakeNet
	Numerics["551"] = "ERR_HOSTUNAVAIL"            //QuakeNet
	Numerics["552"] = "ERR_USINGSLINE"             //QuakeNet
	Numerics["553"] = "ERR_STATSSLINE"             //QuakeNet
	Numerics["600"] = "RPL_LOGON"                  //Bahamut,
	Numerics["601"] = "RPL_LOGOFF"                 //Bahamut,
	Numerics["602"] = "RPL_WATCHOFF"               //Bahamut,
	Numerics["603"] = "RPL_WATCHSTAT"              //Bahamut,
	Numerics["604"] = "RPL_NOWON"                  //Bahamut,
	Numerics["605"] = "RPL_NOWOFF"                 //Bahamut,
	Numerics["606"] = "RPL_WATCHLIST"              //Bahamut,
	Numerics["607"] = "RPL_ENDOFWATCHLIST"         //Bahamut,
	Numerics["608"] = "RPL_WATCHCLEAR"             //Ultimate
	Numerics["610"] = "RPL_MAPMORE"                //Unreal
	Numerics["610"] = "RPL_ISOPER"                 //Ultimate
	Numerics["611"] = "RPL_ISLOCOP"                //Ultimate
	Numerics["612"] = "RPL_ISNOTOPER"              //Ultimate
	Numerics["613"] = "RPL_ENDOFISOPER"            //Ultimate
	Numerics["615"] = "RPL_MAPMORE"                //PTlink
	Numerics["615"] = "RPL_WHOISMODES"             //Ultimate
	Numerics["616"] = "RPL_WHOISHOST"              //Ultimate
	Numerics["617"] = "RPL_DCCSTATUS"              //Bahamut
	Numerics["617"] = "RPL_WHOISBOT"               //Ultimate
	Numerics["618"] = "RPL_DCCLIST"                //Bahamut
	Numerics["619"] = "RPL_ENDOFDCCLIST"           //Bahamut
	Numerics["619"] = "RPL_WHOWASHOST"             //Ultimate
	Numerics["620"] = "RPL_DCCINFO"                //Bahamut
	Numerics["620"] = "RPL_RULESSTART"             //Ultimate
	Numerics["621"] = "RPL_RULES"                  //Ultimate
	Numerics["622"] = "RPL_ENDOFRULES"             //Ultimate
	Numerics["623"] = "RPL_MAPMORE"                //Ultimate
	Numerics["624"] = "RPL_OMOTDSTART"             //Ultimate
	Numerics["625"] = "RPL_OMOTD"                  //Ultimate
	Numerics["626"] = "RPL_ENDOFO"                 //Ultimate
	Numerics["630"] = "RPL_SETTINGS"               //Ultimate
	Numerics["631"] = "RPL_ENDOFSETTINGS"          //Ultimate
	Numerics["640"] = "RPL_DUMPING"                //Unreal
	Numerics["641"] = "RPL_DUMPRPL"                //Unreal
	Numerics["642"] = "RPL_EODUMP"                 //Unreal
	Numerics["660"] = "RPL_TRACEROUTE_HOP"         //KineIRCd
	Numerics["661"] = "RPL_TRACEROUTE_START"       //KineIRCd
	Numerics["662"] = "RPL_MODECHANGEWARN"         //KineIRCd
	Numerics["663"] = "RPL_CHANREDIR"              //KineIRCd
	Numerics["664"] = "RPL_SERVMODEIS"             //KineIRCd
	Numerics["665"] = "RPL_OTHERUMODEIS"           //KineIRCd
	Numerics["666"] = "RPL_ENDOF_GENERIC"          //KineIRCd
	Numerics["670"] = "RPL_WHOWASDETAILS"          //KineIRCd
	Numerics["671"] = "RPL_WHOISSECURE"            //KineIRCd
	Numerics["672"] = "RPL_UNKNOWNMODES"           //Ithildin
	Numerics["673"] = "RPL_CANNOTSETMODES"         //Ithildin
	Numerics["678"] = "RPL_LUSERSTAFF"             //KineIRCd
	Numerics["679"] = "RPL_TIMEONSERVERIS"         //KineIRCd
	Numerics["682"] = "RPL_NETWORKS"               //KineIRCd
	Numerics["687"] = "RPL_YOURLANGUAGEIS"         //KineIRCd
	Numerics["688"] = "RPL_LANGUAGE"               //KineIRCd
	Numerics["689"] = "RPL_WHOISSTAFF"             //KineIRCd
	Numerics["690"] = "RPL_WHOISLANGUAGE"          //KineIRCd
	Numerics["702"] = "RPL_MODLIST"                //RatBox
	Numerics["703"] = "RPL_ENDOFMODLIST"           //RatBox
	Numerics["704"] = "RPL_HELPSTART"              //RatBox
	Numerics["705"] = "RPL_HELPTXT"                //RatBox
	Numerics["706"] = "RPL_ENDOFHELP"              //RatBox
	Numerics["708"] = "RPL_ETRACEFULL"             //RatBox
	Numerics["709"] = "RPL_ETRACE"                 //RatBox
	Numerics["710"] = "RPL_KNOCK"                  //RatBox
	Numerics["711"] = "RPL_KNOCKDLVR"              //RatBox
	Numerics["712"] = "ERR_TOOMANYKNOCK"           //RatBox
	Numerics["713"] = "ERR_CHANOPEN"               //RatBox
	Numerics["714"] = "ERR_KNOCKONCHAN"            //RatBox
	Numerics["715"] = "ERR_KNOCKDISABLED"          //RatBox
	Numerics["716"] = "RPL_TARGUMODEG"             //RatBox
	Numerics["717"] = "RPL_TARGNOTIFY"             //RatBox
	Numerics["718"] = "RPL_UMODEGMSG"              //RatBox
	Numerics["720"] = "RPL_OMOTDSTART"             //RatBox
	Numerics["721"] = "RPL_OMOTD"                  //RatBox
	Numerics["722"] = "RPL_ENDOFOMOTD"             //RatBox
	Numerics["723"] = "ERR_NOPRIVS"                //RatBox
	Numerics["724"] = "RPL_TESTMARK"               //RatBox
	Numerics["725"] = "RPL_TESTLINE"               //RatBox
	Numerics["726"] = "RPL_NOTESTLINE"             //RatBox
	Numerics["771"] = "RPL_XINFO"                  //Ithildin
	Numerics["773"] = "RPL_XINFOSTART"             //Ithildin
	Numerics["774"] = "RPL_XINFOEND"               //Ithildin
	Numerics["972"] = "ERR_CANNOTDOCOMMAND"        //Unreal
	Numerics["973"] = "ERR_CANNOTCHANGEUMODE"      //KineIRCd
	Numerics["974"] = "ERR_CANNOTCHANGECHANMODE"   //KineIRCd
	Numerics["975"] = "ERR_CANNOTCHANGESERVERMODE" //KineIRCd
	Numerics["976"] = "ERR_CANNOTSENDTONICK"       //KineIRCd
	Numerics["977"] = "ERR_UNKNOWNSERVERMODE"      //KineIRCd
	Numerics["979"] = "ERR_SERVERMODELOCK"         //KineIRCd
	Numerics["980"] = "ERR_BADCHARENCODING"        //KineIRCd
	Numerics["981"] = "ERR_TOOMANYLANGUAGES"       //KineIRCd
	Numerics["982"] = "ERR_NOLANGUAGE"             //KineIRCd
	Numerics["983"] = "ERR_TEXTTOOSHORT"           //KineIRCd
	Numerics["999"] = "ERR_NUMERIC_ERR"            //Bahamut
}
