package gaf

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestGetSidelined(t *testing.T) {
	server := httptest.NewServer(getHandler(t, &HandlerHelper{
		ResponseBody: `{"get":"sidelined","parameters":{"player":"276"},"errors":[],"results":27,"paging":{"current":1,"total":1},"response":[{"type":"Suspended","start":"2020-02-26","end":"2020-03-01"},{"type":"Hip/Thigh Injury","start":"2020-02-02","end":"2020-02-10"},{"type":"Groin/Pelvis Injury","start":"2019-10-11","end":"2019-11-20"},{"type":"Ankle/Foot Injury","start":"2019-08-01","end":"2019-08-23"},{"type":"Suspended","start":"2019-05-15","end":"2019-05-27"},{"type":"Ankle/Foot Injury","start":"2019-01-24","end":"2019-04-20"},{"type":"Groin Strain","start":"2018-12-03","end":"2019-01-02"},{"type":"Groin Strain","start":"2018-11-21","end":"2018-11-27"},{"type":"Broken Toe","start":"2018-02-26","end":"2018-05-20"},{"type":"Thigh Muscle Strain","start":"2018-01-20","end":"2018-01-25"},{"type":"Rib Injury","start":"2018-01-11","end":"2018-01-16"},{"type":"Suspended","start":"2017-12-05","end":"2017-12-11"},{"type":"Thigh Muscle Strain","start":"2017-11-03","end":"2017-11-15"},{"type":"Suspended","start":"2017-10-23","end":"2017-10-28"},{"type":"Ankle/Foot Injury","start":"2017-09-21","end":"2017-09-25"},{"type":"Suspended","start":"2017-04-09","end":"2017-04-27"},{"type":"Suspended","start":"2016-12-04","end":"2016-12-11"},{"type":"Suspended","start":"2016-03-04","end":"2016-03-07"},{"type":"Hamstring","start":"2016-01-21","end":"2016-01-26"},{"type":"Hamstring","start":"2015-12-08","end":"2015-12-16"},{"type":"Virus","start":"2015-08-09","end":"2015-08-26"},{"type":"Suspended","start":"2015-03-01","end":"2015-03-09"},{"type":"Sprained Ankle","start":"2014-08-22","end":"2014-08-29"},{"type":"Vertebral Fracture","start":"2014-07-05","end":"2014-08-05"},{"type":"Ankle/Foot Injury","start":"2014-04-17","end":"2014-05-10"},{"type":"Sprained Ankle","start":"2014-01-17","end":"2014-02-14"},{"type":"Suspended","start":"2013-12-15","end":"2013-12-23"}]}`,
	}))
	defer server.Close()

	client := NewTestClientWithCustomHandler(t, server)
	options := &SidelinedOptions{Player: 276}

	res, err := client.GetSidelined(context.Background(), options)
	require.NoError(t, err)

	require.Len(t, *res, 27)
	require.Equal(t, "Suspended", (*res)[0].Type)
	require.Equal(t, "2020-02-26", (*res)[0].Start)
	require.Equal(t, "2020-03-01", (*res)[0].End)
	require.Equal(t, "Suspended", (*res)[26].Type)
	require.Equal(t, "2013-12-15", (*res)[26].Start)
	require.Equal(t, "2013-12-23", (*res)[26].End)

}
