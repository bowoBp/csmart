package bookingbook

type ResponseBookLibrary struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	SubjectType string `json:"subject_type"`
	WorkCount   int    `json:"work_count"`
	Works       []struct {
		Key               string   `json:"key"`
		Title             string   `json:"title"`
		EditionCount      int      `json:"edition_count"`
		CoverID           int      `json:"cover_id"`
		CoverEditionKey   string   `json:"cover_edition_key"`
		Subject           []string `json:"subject"`
		IaCollection      []string `json:"ia_collection"`
		Lendinglibrary    bool     `json:"lendinglibrary"`
		Printdisabled     bool     `json:"printdisabled"`
		LendingEdition    string   `json:"lending_edition"`
		LendingIdentifier string   `json:"lending_identifier"`
		Authors           []struct {
			Key  string `json:"key"`
			Name string `json:"name"`
		} `json:"authors"`
		FirstPublishYear int    `json:"first_publish_year"`
		Ia               string `json:"ia"`
		PublicScan       bool   `json:"public_scan"`
		HasFulltext      bool   `json:"has_fulltext"`
		Availability     struct {
			Status              string      `json:"status"`
			AvailableToBrowse   interface{} `json:"available_to_browse"`
			AvailableToBorrow   interface{} `json:"available_to_borrow"`
			AvailableToWaitlist interface{} `json:"available_to_waitlist"`
			IsPrintdisabled     interface{} `json:"is_printdisabled"`
			IsReadable          interface{} `json:"is_readable"`
			IsLendable          interface{} `json:"is_lendable"`
			IsPreviewable       bool        `json:"is_previewable"`
			Identifier          string      `json:"identifier"`
			Isbn                interface{} `json:"isbn"`
			Oclc                interface{} `json:"oclc"`
			OpenlibraryWork     interface{} `json:"openlibrary_work"`
			OpenlibraryEdition  interface{} `json:"openlibrary_edition"`
			LastLoanDate        interface{} `json:"last_loan_date"`
			NumWaitlist         interface{} `json:"num_waitlist"`
			LastWaitlistDate    interface{} `json:"last_waitlist_date"`
			IsRestricted        bool        `json:"is_restricted"`
			IsBrowseable        interface{} `json:"is_browseable"`
			Src                 string      `json:"__src__"`
		} `json:"availability"`
	} `json:"works"`
}
