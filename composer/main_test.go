package main

import (
	"fmt"
	"testing"
)

func TestCompose(t *testing.T) {
	// trainingSongs := []string{"fDfDfadsptupauOasufDfDfadsptupausapasdfogfdifdsudsaufuffxDfDfDfDfDfDfadsptupauOasufDfDfadsptupausap", "9eyieyieyieyi8eyieyieyieyiEyiEyiEYoEYo6eToeyi6eyuwTu9qeyeyieyipeyppuoeuoeuopuop9pyieyiPyoEyo8ptieti8otusuo4itietietietiWtiWtiWtiOtiOOtIOtIOtIOtIOTiWTiWyioyi", "fsappaspasfgfddapOOpaOpadgDflkjhgdkjhhfsdffdfgdsfdssfh", "tuosfosftuosfosftypdgpdgtypdgpdgryodgodgryodgodgtuosfosftuosfosftupfjpfjtupfjpfjtyIpdIpdtyIpdIpdryodhodhryodhodhrtuosuosrtuosuosetuosuosetuosuos9eyIsyIs9eyIsyIswryoayoawryoayoawEuoSuoSwEuoSuoSqeypdypdqeypdypdqwyiaqwyia0wtostos0wtostos0qetieti0qetieti9qetieti9qetieti59wriwri59wriwri80wtuwtu80wtuwtu8wetuetu8wetuetu4qetuetu4qetuetu48etyety48etyety5qrtyrty5qrtyrty5qwrywry5qwrywry50wtuwtu50wtuwtu59wtiwti59wtiwti59wriwri59wriwri59etieti59etieti50wtowto50wto59wtiwti59wtiwti59wriwri59wri18weuweu18weuweu18qetiteteqe9q917oadgdadaoayiuy18uos", "fulfululkjhjkhfukfukfkffulfululkjhjkhfukfukfkfstjstjtjhgfghfdyjdyjdjdfulfululkjhjkhfukfukfkffulfululkjhjkhfukfukfkf"}
	trainingSongs := loadData("../data.csv")
	orderedStatisticsMap, cumulativeArrayMap := constructStats(trainingSongs)

	length := 110
	currentChar := "a"
	song := compose(currentChar, length, orderedStatisticsMap, cumulativeArrayMap)

	if len(song) != 110 {
		fmt.Printf("len should be 110,but is %d", len(song))
		t.Fail()
	}
}
