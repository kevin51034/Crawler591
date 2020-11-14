package main


type HouseList struct {
	Status int `json:"status"`
	Data   struct {
		TopData []struct {
			IsNewIndex int    `json:"is_new_index"`
			IsNewList  int    `json:"is_new_list"`
			Type       int    `json:"type"`
			PostID     int    `json:"post_id"`
			IsAd       int    `json:"isAd"`
			Regionid   int    `json:"regionid"`
			PhotoNum   int    `json:"photoNum"`
			ClassLast  string `json:"classLast"`
			DetailURL  string `json:"detail_url"`
			Address    string `json:"address"`
			ImgSrc     string `json:"img_src"`
			Alt        string `json:"alt"`
			Address2   string `json:"address_2"`
			SectionStr string `json:"section_str"`
			Region     int    `json:"region"`
			KindStr    string `json:"kind_str"`
			Area       string `json:"area"`
			PriceStr   string `json:"price_str"`
			RecomNum   int    `json:"recom_num"`
			Address3   string `json:"address_3"`
			Ico        string `json:"ico"`
			Price      string `json:"price"`
			PriceUnit  string `json:"price_unit"`
			OnepxImg   string `json:"onepxImg"`
		} `json:"topData"`
		Biddings []interface{} `json:"biddings"`
		Data     []struct {
			ID                  int         `json:"id"`
			UserID              int         `json:"user_id"`
			Address             string      `json:"address"`
			Type                string      `json:"type"`
			PostID              int         `json:"post_id"`
			Regionid            int         `json:"regionid"`
			Sectionid           int         `json:"sectionid"`
			Streetid            int         `json:"streetid"`
			Room                int         `json:"room"`
			Area                float64     `json:"area"`
			Price               string      `json:"price"`
			Storeprice          int         `json:"storeprice"`
			CommentTotal        int         `json:"comment_total"`
			CommentUnread       int         `json:"comment_unread"`
			CommentLtime        int         `json:"comment_ltime"`
			Hasimg              int         `json:"hasimg"`
			Kind                int         `json:"kind"`
			Shape               int         `json:"shape"`
			Houseage            int         `json:"houseage"`
			Posttime            string      `json:"posttime"`
			Updatetime          int         `json:"updatetime"`
			Refreshtime         int         `json:"refreshtime"`
			Checkstatus         int         `json:"checkstatus"`
			Status              string      `json:"status"`
			Closed              int         `json:"closed"`
			Living              string      `json:"living"`
			Condition           string      `json:"condition"`
			Isvip               int         `json:"isvip"`
			Mvip                int         `json:"mvip"`
			IsCombine           int         `json:"is_combine"`
			Cover               string      `json:"cover"`
			Browsenum           int         `json:"browsenum"`
			BrowsenumAll        int         `json:"browsenum_all"`
			Floor2              int         `json:"floor2"`
			Floor               int         `json:"floor"`
			Ltime               string      `json:"ltime"`
			CasesID             int         `json:"cases_id"`
			SocialHouse         int         `json:"social_house"`
			Distance            int         `json:"distance"`
			SearchName          string      `json:"search_name"`
			Mainarea            interface{} `json:"mainarea"`
			BalconyArea         interface{} `json:"balcony_area"`
			Groundarea          interface{} `json:"groundarea"`
			Linkman             string      `json:"linkman"`
			Housetype           int         `json:"housetype"`
			StreetName          string      `json:"street_name"`
			AlleyName           string      `json:"alley_name"`
			LaneName            string      `json:"lane_name"`
			AddrNumberName      string      `json:"addr_number_name"`
			KindNameImg         string      `json:"kind_name_img"`
			AddressImg          string      `json:"address_img"`
			CasesName           string      `json:"cases_name"`
			Layout              string      `json:"layout"`
			LayoutStr           string      `json:"layout_str"`
			Allfloor            int         `json:"allfloor"`
			FloorInfo           string      `json:"floorInfo"`
			HouseImg            string      `json:"house_img"`
			Houseimg            interface{} `json:"houseimg"`
			Cartplace           string      `json:"cartplace"`
			SpaceTypeStr        string      `json:"space_type_str"`
			PhotoAlt            string      `json:"photo_alt"`
			Addition4           int         `json:"addition4"`
			Addition2           int         `json:"addition2"`
			Addition3           int         `json:"addition3"`
			Vipimg              string      `json:"vipimg"`
			Vipstyle            string      `json:"vipstyle"`
			VipBorder           string      `json:"vipBorder"`
			NewListCommentTotal int         `json:"new_list_comment_total"`
			CommentClass        string      `json:"comment_class"`
			PriceHide           string      `json:"price_hide"`
			KindName            string      `json:"kind_name"`
			PhotoNum            string      `json:"photoNum"`
			Filename            string      `json:"filename"`
			NickName            string      `json:"nick_name"`
			NewImg              string      `json:"new_img"`
			Regionname          string      `json:"regionname"`
			Sectionname         string      `json:"sectionname"`
			IconName            string      `json:"icon_name"`
			IconClass           string      `json:"icon_class"`
			Fulladdress         string      `json:"fulladdress"`
			AddressImgTitle     string      `json:"address_img_title"`
			BrowsenumName       string      `json:"browsenum_name"`
			Unit                string      `json:"unit"`
			Houseid             int         `json:"houseid"`
			RegionName          string      `json:"region_name"`
			SectionName         string      `json:"section_name"`
			AddInfo             string      `json:"addInfo"`
			OnepxImg            string      `json:"onepxImg"`
		} `json:"data"`
		Page string `json:"page"`
	} `json:"data"`
	Records          string        `json:"records"`
	IsRecom          int           `json:"is_recom"`
	DealRecom        []interface{} `json:"deal_recom"`
	OnlineSocialUser int           `json:"online_social_user"`
	BluekaiData      struct {
		RegionID         string `json:"region_id"`
		SectionID        string `json:"section_id"`
		SalePrice        int    `json:"sale_price"`
		RentalPrice      int    `json:"rental_price"`
		UnitPricePerPing string `json:"unit_price_per_ping"`
		Room             string `json:"room"`
		Shape            string `json:"shape"`
		MrtCity          string `json:"mrt_city"`
		MrtLine          string `json:"mrt_line"`
		Tag              int    `json:"tag"`
		Type             string `json:"type"`
		Kind             string `json:"kind"`
		Page             string `json:"page"`
	} `json:"bluekai_data"`
}