package resolution_cli

import (
	"os"
	"path"
	"testing"

	"github.com/rendon/testcli"
	"github.com/stretchr/testify/assert"
)

func commandPath() string {
	var pwd, _ = os.Getwd()
	return path.Join(pwd, "test-cli")
}

func assertCommandResult(t *testing.T, cmd *testcli.Cmd, expectedOutput string) {
	assert.Equal(t, expectedOutput, cmd.Stdout())
	assert.Nil(t, cmd.Error())
	assert.Empty(t, cmd.Stderr())
}

func TestCliResolve(t *testing.T) {
	t.Parallel()
	cmd := testcli.Command(commandPath(), "resolve", "-d", "brad.crypto")
	cmd.Run()
	expectedOutput := `{
   "records": {
      "crypto.BTC.address": "bc1q359khn0phg58xgezyqsuuaha28zkwx047c0c3y",
      "crypto.ETH.address": "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8",
      "gundb.public_key.value": "pqeBHabDQdCHhbdivgNEc74QO-x8CPGXq4PKWgfIzhY.7WJR5cZFuSyh1bFwx0GWzjmrim0T5Y6Bp0SSK0im3nI",
      "gundb.username.value": "0x8912623832e174f2eb1f59cc3b587444d619376ad5bf10070e937e0dc22b9ffb2e3ae059e6ebf729f87746b2f71e5d88ec99c1fb3c7c49b8617e2520d474c48e1c",
      "ipfs.html.value": "QmdyBw5oTgCtTLQ18PbDvPL8iaLoEPhSyzD91q9XmgmAjb",
      "ipfs.redirect_domain.value": "https://abbfe6z95qov3d40hf6j30g7auo7afhp.mypinata.cloud/ipfs/Qme54oEzRkgooJbCDr78vzKAWcv6DDEZqRhhDyDtzgrZP6"
   }
}
`
	assertCommandResult(t, cmd, expectedOutput)
}

func TestCliAddr(t *testing.T) {
	t.Parallel()
	cmd := testcli.Command(commandPath(), "resolve", "addr", "ETH", "-d", "brad.crypto")
	cmd.Run()
	expectedOutput := `"0x8aaD44321A86b170879d7A244c1e8d360c99DdA8"
`
	assertCommandResult(t, cmd, expectedOutput)
}

func TestCliAddrVersion(t *testing.T) {
	t.Parallel()
	cmd := testcli.Command(commandPath(), "resolve", "addr-version", "-v", "ERC20", "-c", "USDT", "-d", "udtestdev-usdt.crypto")
	cmd.Run()
	expectedOutput := `"0xe7474D07fD2FA286e7e0aa23cd107F8379085037"
`
	assertCommandResult(t, cmd, expectedOutput)
}

func TestCliDns(t *testing.T) {
	t.Parallel()
	cmd := testcli.Command(commandPath(), "resolve", "dns", "CNAME", "-d", "udtestdev-dns-cname.crypto")
	cmd.Run()
	expectedOutput := `[
   {
      "Type": "CNAME",
      "TTL": 1111,
      "Value": "example.com."
   }
]
`
	assertCommandResult(t, cmd, expectedOutput)
}

func TestCliEmail(t *testing.T) {
	t.Parallel()
	cmd := testcli.Command(commandPath(), "resolve", "email", "-d", "reseller-test-paul019.crypto")
	cmd.Run()
	expectedOutput := `"paul@unstoppabledomains.com"
`
	assertCommandResult(t, cmd, expectedOutput)
}

func TestCliHttpUrl(t *testing.T) {
	t.Parallel()
	cmd := testcli.Command(commandPath(), "resolve", "http-url", "-d", "brad.crypto")
	cmd.Run()
	expectedOutput := `"https://abbfe6z95qov3d40hf6j30g7auo7afhp.mypinata.cloud/ipfs/Qme54oEzRkgooJbCDr78vzKAWcv6DDEZqRhhDyDtzgrZP6"
`
	assertCommandResult(t, cmd, expectedOutput)
}

func TestCliIpfsHash(t *testing.T) {
	t.Parallel()
	cmd := testcli.Command(commandPath(), "resolve", "ipfs-hash", "-d", "brad.crypto")
	cmd.Run()
	expectedOutput := `"QmdyBw5oTgCtTLQ18PbDvPL8iaLoEPhSyzD91q9XmgmAjb"
`
	assertCommandResult(t, cmd, expectedOutput)
}

func TestCliOwner(t *testing.T) {
	t.Parallel()
	cmd := testcli.Command(commandPath(), "resolve", "owner", "-d", "brad.crypto")
	cmd.Run()
	expectedOutput := `"0x8aaD44321A86b170879d7A244c1e8d360c99DdA8"
`
	assertCommandResult(t, cmd, expectedOutput)
}

func TestCliResolver(t *testing.T) {
	t.Parallel()
	cmd := testcli.Command(commandPath(), "resolve", "resolver", "-d", "brad.crypto")
	cmd.Run()
	expectedOutput := `"0xb66DcE2DA6afAAa98F2013446dBCB0f4B0ab2842"
`
	assertCommandResult(t, cmd, expectedOutput)
}

func TestCliRecords(t *testing.T) {
	t.Parallel()
	cmd := testcli.Command(commandPath(), "resolve", "records", "crypto.ETH.address", "crypto.BTC.address", "-d", "brad.crypto")
	cmd.Run()
	expectedOutput := `{
   "records": {
      "crypto.BTC.address": "bc1q359khn0phg58xgezyqsuuaha28zkwx047c0c3y",
      "crypto.ETH.address": "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8"
   }
}
`
	assertCommandResult(t, cmd, expectedOutput)
}

func TestCliZilResolve(t *testing.T) {
	t.Parallel()
	cmd := testcli.Command(commandPath(), "resolve", "-d", "brad.zil")
	cmd.Run()
	expectedOutput := `{
   "records": {
      "crypto.BCH.address": "qrq4sk49ayvepqz7j7ep8x4km2qp8lauvcnzhveyu6",
      "crypto.BTC.address": "1EVt92qQnaLDcmVFtHivRJaunG2mf2C3mB",
      "crypto.DASH.address": "XnixreEBqFuSLnDSLNbfqMH1GsZk7cgW4j",
      "crypto.ETH.address": "0x45b31e01AA6f42F0549aD482BE81635ED3149abb",
      "crypto.LTC.address": "LetmswTW3b7dgJ46mXuiXMUY17XbK29UmL",
      "crypto.XMR.address": "447d7TVFkoQ57k3jm3wGKoEAkfEym59mK96Xw5yWamDNFGaLKW5wL2qK5RMTDKGSvYfQYVN7dLSrLdkwtKH3hwbSCQCu26d",
      "crypto.ZEC.address": "t1h7ttmQvWCSH1wfrcmvT4mZJfGw2DgCSqV",
      "crypto.ZIL.address": "zil1yu5u4hegy9v3xgluweg4en54zm8f8auwxu0xxj",
      "ipfs.html.value": "QmVaAtQbi3EtsfpKoLzALm6vXphdi2KjMgxEDKeGg6wHuK",
      "ipfs.redirect_domain.value": "www.unstoppabledomains.com"
   }
}
`
	assertCommandResult(t, cmd, expectedOutput)
}
