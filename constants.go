package afreecachat

// Reference : https://github.com/wakscord/afreeca
const (
	SVC_KEEPALIVE             = 0 /* pingpong */
	SVC_LOGIN                 = 1 /* 로그인 핸드쉐이크 */
	SVC_JOINCH                = 2 /* 채널 입장 핸드쉐이크 */
	SVC_QUITCH                = 3 /* 강제 퇴장 */
	SVC_CHUSER                = 4 /* 입장/퇴장 */
	SVC_CHATMESG              = 5 /* 채팅 메시지 */
	SVC_SETCHNAME             = 6
	SVC_SETBJSTAT             = 7
	SVC_SETDUMB               = 8 /* 채금 */
	SVC_DIRECTCHAT            = 9
	SVC_NOTICE                = 10 /* 사용하지 않음 */
	SVC_KICK                  = 11 /* 사용하지 않음 */
	SVC_SETUSERFLAG           = 12
	SVC_SETSUBBJ              = 13
	SVC_SETNICKNAME           = 14
	SVC_SVRSTAT               = 15 /* 사용하지 않음 */
	SVC_NULL_16               = 16 /* 사용하지 않음 */
	SVC_CLUBCOLOR             = 17
	SVC_SENDBALLOON           = 18 /* 별풍선 */
	SVC_ICEMODE               = 19 /* 얼음 */
	SVC_SENDFANLETRTRER       = 20
	SVC_ICEMODE_EX            = 21 /* 얼음 */ // log상 21, 19 순서임.
	SVC_GET_ICEMODE_RELAY     = 22 /* 사용하지 않음 */
	SVC_SLOWMODE              = 23 /* 슬로우 모드 */
	SVC_RELOADBURNLEVEL       = 24 /* 사용하지 않음 */
	SVC_BLINDKICK             = 25 /* 사용하지 않음 */
	SVC_MANAGERCHAT           = 26 /* 매니저 채팅, 읽을 수 있는 권리는 매니저 flag 이상만 */
	SVC_APPENDDATA            = 27 /* 사용하지 않음 */
	SVC_BASEBALLEVENT         = 28 /* 사용하지 않음 */
	SVC_PAIDITEM              = 29 /* 사용하지 않음 */
	SVC_TOPFAN                = 30 /* 사용하지 않음 */ /* 열혈? */
	SVC_SNSMESSAGE            = 31 /* 사용하지 않음 */
	SVC_SNSMODE               = 32 /* 사용하지 않음 */
	SVC_SENDBALLOONSUB        = 33
	SVC_SENDFANLETRTRERSUB    = 34
	SVC_TOPFANSUB             = 35 /* 사용하지 않음 */
	SVC_BJSTICKERITEM         = 36 /* 사용하지 않음 */
	SVC_CHOCOLATE             = 37
	SVC_CHOCOLATESUB          = 38
	SVC_TOPCLAN               = 39 /* 사용하지 않음 */
	SVC_TOPCLANSUB            = 40 /* 사용하지 않음 */
	SVC_SUPERCHAT             = 41 /* 사용하지 않음 */
	SVC_UPDATETICKET          = 42 /* 사용하지 않음 */
	SVC_NOTIGAMERANKER        = 43 /* 사용하지 않음 */
	SVC_STARCOIN              = 44
	SVC_SENDQUICKVIEW         = 45 /* 퀵 뷰 선물 */
	SVC_ITEMSTATUS            = 46 /* 사용하지 않음 */
	SVC_ITEMUSING             = 47
	SVC_USEQUICKVIEW          = 48
	SVC_NOTIFY_POLL           = 50 /* 투표 */
	SVC_CHATBLOCKMODE         = 51 /* 사용하지 않음 */
	SVC_BDM_ADDBLACKINFO      = 52 /* 블랙리스트..? */
	SVC_SETBROADINFO          = 53 /* 사용하지 않음 */
	SVC_BAN_WORD              = 54
	SVC_SENDADMINNOTICE       = 58 /* 어드민 메시지 */
	SVC_FREECAT_OWNER_JOIN    = 65
	SVC_BUYGOODS              = 70
	SVC_BUYGOODSSUB           = 71
	SVC_SENDPROMOTION         = 72 /* 사용하지 않음 */
	SVC_NOTIFY_VR             = 74
	SVC_NOTIFY_MOBBROAD_PAUSE = 75
	SVC_KICK_AND_CANCEL       = 76
	SVC_KICK_USERLIST         = 77
	SVC_ADMIN_CHUSER          = 78
	SVC_CLIDOBAEINFO          = 79
	SVC_VOD_BALLOON           = 86
	SVC_ADCON_EFFECT          = 87
	SVC_SVC_KICK_MSG_STATE    = 90
	SVC_FOLLOW_ITEM           = 91 /* 신규 구독 */
	SVC_ITEM_SELL_EFFECT      = 92
	SVC_FOLLOW_ITEM_EFFECT    = 93 /* 연속 구독 */
	SVC_TRANSLATION_STATE     = 94
	SVC_TRANSLATION           = 95
	SVC_GIFT_TICKET           = 102
	SVC_VODADCON              = 103
	SVC_BJ_NOTICE             = 104 /* BJ 공지 */
	SVC_VIDEOBALLOON          = 105
	SVC_STATION_ADCON         = 107
	SVC_SENDSUBSCRIPTION      = 108 /* 구독권 선물 */
	SVC_OGQ_EMOTICON          = 109
	SVC_ITEM_DROPS            = 111
	SVC_VIDEOBALLOON_LINK     = 117 /* 사용하지 않음 */
	SVC_OGQ_EMOTICON_GIFT     = 118 /* OGQ 이모티콘 선물 */
	SVC_AD_IN_BROAD_JSON      = 119
	SVC_GEM_ITEMSEND          = 120
	SVC_MISSION               = 121 /* 도전 미션 */
	SVC_LIVE_CAPTION          = 122
	SVC_MISSION_SETTLE        = 125
	SVC_SET_ADMIN_FLAG        = 126
	SVC_CHUSER_EXTEND         = 127 /* 구독자 리스트 */
	SVC_ADMIN_CHUSER_EXTEND   = 128
)
