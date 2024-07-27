package databatch

import (
	"context"
	"fmt"
	"job-seek/pkg/protos"
	"job-seek/services/job_search_service/server"
	"time"

	"github.com/sirupsen/logrus"
)

func JobDumping(serviceProxy *server.JobSearchServiceServerImpl, log *logrus.Logger) {
	timeStart := time.Now()
	fmt.Println("Job Dumping ")
	classificationList := []int32{
		1200, 6140, 1303, 6141, 6143, 6142, 6144, 6145, 6146, 6147, 6148, 6149, 6150, 6151, 6152, 6153, 6154, 6155, 6156, 6157,
		1468, 6158, 6159, 6160, 6161, 6162, 6251, 6252, 6253, 6254, 6255, 6256, 6257, 6258, 6259, 6260, 6304, 6305, 6306, 6307, 6308,
		6309, 6310, 6311, 6312, 6313, 6314, 6315, 6316, 1203, 6174, 6175, 6177, 6178, 6176, 1352, 6179, 6180, 6181, 1392, 1404, 6182,
		6183, 6184, 6185, 6186, 6187, 1204, 6084, 6085, 6086, 6087, 6088, 6089, 6090, 6091, 7019, 7020, 7021, 7022, 7023, 7024, 6163,
		6164, 6165, 6166, 6167, 6168, 6169, 6170, 6171, 6172, 6173, 1206, 6113, 1387, 6114, 6115, 6116, 6117, 6118, 6119, 6120, 6121,
		6122, 6076, 6077, 6078, 6079, 6080, 6081, 6082, 6083, 6263, 6265, 6264, 6267, 6268, 6266, 6269, 6270, 6271, 6272, 6273, 6274,
		6123, 6124, 6125, 6126, 6127, 6128, 6129, 6130, 6132, 6133, 6134, 6135, 6136, 6131, 6137, 6138, 6139, 1209, 6022, 6023, 6024,
		6026, 6027, 6028, 6025, 6029, 6030, 6031, 6032, 6033, 6034, 6035, 6036, 6038, 6037, 6039, 6040, 6041, 6042, 6205, 6206, 6207,
		6208, 6209, 6210, 6211, 6212, 6213, 6214, 1210, 1314, 1322, 1378, 1409, 1410, 1411, 1450, 6359, 6360, 6361, 1211, 6329, 6330,
		1353, 1372, 6331, 1386, 6332, 6333, 6334, 6335, 6336, 6337, 6338, 6339, 6340, 6341, 6342, 6343, 6344, 6345, 6346, 6347, 6348,
		6349, 1467, 6350, 1470, 6352, 6351, 6370, 6353, 6354, 6355, 1212, 1315, 1332, 6052, 6053, 1405, 1415, 6054, 6055, 6056, 1537,
		1542, 1549, 6057, 6317, 6281, 1214, 1216, 6092, 6008, 6058, 1220, 6043, 6362, 1223, 6261, 6246, 1225, 1313, 1328, 6226, 6227,
		6228, 1345, 6229, 6230, 6231, 6232, 1406, 6233, 6234, 6235, 6236, 6237, 6238, 6240, 6239, 6241, 6242, 6243, 6244, 6245,
	}
	// workLocale := "Sydney"
	allowMixCache := true
	salaryType := protos.SalaryType_ANNUAL

	for _, classification := range classificationList {
		req := &protos.JobSearchRequest{
			// UserId:     "32130c50-7afc-4130-9491-4caddc30f81d",
			SalaryType:     &salaryType,
			Keywords:       []string{""},
			Classification: &classification,
			// WorkLocale:     &workLocale,
			AllowMixCache: &allowMixCache,
		}
		fmt.Println("Job Dumping ", classification)
		testCtx := context.Background()
		resp, err := serviceProxy.JobSearch(testCtx, req)
		testCtx.Done()
		if err != nil {
			log.WithFields(map[string]interface{}{
				"model":    "JobSearchServiceServerImpl",
				"error":    err,
				"request":  req,
				"response": resp,
				"method":   "JobSearch",
			}).Error("fail to dumping jobs into database")
			fmt.Println("JobDumping Error")
		}
		fmt.Println("Job Dumping end", classification)
		time.Sleep(15 * time.Second)

	}

	timeTook := time.Since(timeStart)
	fmt.Printf("Time took : %f secs \n", timeTook.Seconds())
	fmt.Println("Cooling down for 10 minutes")
	time.Sleep(10 * time.Minute)
	fmt.Println("TestJobSearch Done")
}
