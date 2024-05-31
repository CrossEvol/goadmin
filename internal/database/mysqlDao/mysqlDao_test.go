package mysqlDao

import (
	"context"
	"fmt"
	"github.com/crossevol/goadmin/assets"
	"github.com/crossevol/goadmin/internal/database"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"io/fs"
	"os"
	"regexp"
	"strings"
	"testing"
)

type MySQLDaoTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
	db                            *database.DB
	queries                       *Queries
	dbDsn                         string
	dbName                        string
	dbUrl                         string
}

func (suite *MySQLDaoTestSuite) SetupSuite() {
	envBytes, err := assets.EmbeddedFiles.ReadFile("env/mysqlDaoTest")
	if err != nil {
		fmt.Println("Failed to read .env from assets/env ...")
	}
	err = os.WriteFile(".env", envBytes, fs.ModePerm)
	if err != nil {
		fmt.Println("Failed to write temp .env...")
	}

	godotenv.Load(".env")
	suite.dbDsn = os.Getenv("DB_DSN")
	suite.dbDsn, suite.dbUrl, suite.dbName = processDbDsn(suite.dbDsn)
	suite.db, err = database.New(suite.dbUrl, false)
	if err != nil {
		fmt.Println(err)
	}
	suite.db.Exec(fmt.Sprintf("CREATE DATABASE %s CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci", suite.dbName))
	suite.db, err = database.New(suite.dbDsn, true)
	suite.queries = New(suite.db)
}

func (suite *MySQLDaoTestSuite) TearDownSuite() {
	if suite.db != nil {
		suite.db.Exec(fmt.Sprintf("DROP DATABASE %s", suite.dbName))
		suite.db.Close()
		suite.queries.Close()
	}
	os.Remove(".env")
}

func (suite *MySQLDaoTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5
}

func (suite *MySQLDaoTestSuite) TearDownTest() {
}

func (suite *MySQLDaoTestSuite) TestMySQLDao() {
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
	ctx := context.Background()

	_, err := suite.queries.GetGoadminMenus(ctx)
	require.Nil(suite.T(), err)
	_, err = suite.queries.GetGoadminOperationLogs(ctx)
	require.Nil(suite.T(), err)
	_, err = suite.queries.GetGoadminPermissions(ctx)
	require.Nil(suite.T(), err)
	_, err = suite.queries.GetGoadminRoleMenus(ctx)
	require.Nil(suite.T(), err)
	_, err = suite.queries.GetGoadminRolePermissions(ctx)
	require.Nil(suite.T(), err)
	_, err = suite.queries.GetGoadminRoleUsers(ctx)
	require.Nil(suite.T(), err)
	_, err = suite.queries.GetGoadminRoles(ctx)
	require.Nil(suite.T(), err)
	_, err = suite.queries.GetGoadminSessions(ctx)
	require.Nil(suite.T(), err)
	_, err = suite.queries.GetGoadminSites(ctx)
	require.Nil(suite.T(), err)
	_, err = suite.queries.CountGoadminUserPermissions(ctx)
	require.Nil(suite.T(), err)
	_, err = suite.queries.GetGoadminUsers(ctx)
	require.Nil(suite.T(), err)
}

func TestMySQLDaoTestSuite(t *testing.T) {
	suite.Run(t, new(MySQLDaoTestSuite))
}

func processDbDsn(srcDbDsn string) (string, string, string) {
	dbUrl := srcDbDsn[:strings.Index(srcDbDsn, "/")+1]
	// Regular expression to match the string between "/" and "?"
	re := regexp.MustCompile(`/([^/?]*)\?`)

	// Find and replace with random string
	dbName := re.FindStringSubmatch(srcDbDsn)[1] + uuid.New().String()

	destDbDsn := re.ReplaceAllString(srcDbDsn, "/"+dbName+"?")
	destDbDsn = strings.ReplaceAll(destDbDsn, "-", "_")
	dbName = strings.ReplaceAll(dbName, "-", "_")

	return destDbDsn, dbUrl, dbName
}
