package agent

import "testing"

func TestSafari6(t *testing.T) {
	s := "Mozilla/5.0 (iPad; CPU OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5355d Safari/8536.25"
	testAgent(t, s, "6.0", "Safari")
}

func TestSafari517(t *testing.T) {
	s := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/537.13+ (KHTML, like Gecko) Version/5.1.7 Safari/534.57.2"
	testAgent(t, s, "5.1.7", "Safari")
}

func TestSafari504(t *testing.T) {
	s := "Mozilla/5.0 (Windows; U; Windows NT 6.0; en-US) AppleWebKit/533.20.25 (KHTML, like Gecko) Version/5.0.4 Safari/533.20.27"
	testAgent(t, s, "5.0.4", "Safari")
}

func TestChrome29(t *testing.T) {
	s := "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/29.0.1547.2 Safari/537.36"
	testAgent(t, s, "29.0.1547.2", "Chrome")
}

func TestChrome27(t *testing.T) {
	s := "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.93 Safari/537.36"
	testAgent(t, s, "27.0.1453.93", "Chrome")
}

func TestFirefox24(t *testing.T) {
	s := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:24.0) Gecko/20100101 Firefox/24.0"
	testAgent(t, s, "24.0", "Firefox")
}

func TestFirefox23(t *testing.T) {
	s := "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:23.0) Gecko/20131011 Firefox/23.0"
	testAgent(t, s, "23.0", "Firefox")
}

func testAgent(t *testing.T, s, v, n string) {
	a := Parse(s)

	if a.Version != v {
		t.Errorf("version should be %q, but was %q", v, a.Version)
	}

	if a.Name != n {
		t.Errorf("name should be %q, but was %q", n, a.Name)
	}
}

/*
// firefox
Mozilla/5.0 (Windows NT 6.1; rv:22.0) Gecko/20130405 Firefox/22.0

// ie
Mozilla/5.0 (compatible; MSIE 10.6; Windows NT 6.1; Trident/5.0; InfoPath.2; SLCC1; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET CLR 2.0.50727) 3gpp-gba UNTRUSTED/1.0
Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; WOW64; Trident/6.0)
Mozilla/5.0 (compatible; MSIE 10.0; Macintosh; Intel Mac OS X 10_7_3; Trident/6.0)
Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; .NET CLR 1.1.4322; .NET CLR 2.0.50727)

// mobile
Mozilla/5.0 (PLAYSTATION 3; 3.55)
curl/7.9.8 (i686-pc-linux-gnu) libcurl 7.9.8 (OpenSSL 0.9.6b) (ipv6 enabled)

*/
