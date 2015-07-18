package mgnzip

import (
	"testing"
	"github.com/stretchr/testify/suite"
)

type ZipTestSuite struct {
	suite.Suite
	OutputDir string
}

func TestZipTestSuite(t *testing.T) {
	sample := new(ZipTestSuite)
	suite.Run(t, sample)
	//assert.Equal(t, sample.VariableThatShouldStartAtFive, 4)
}

func (suite *ZipTestSuite) SetupTest() {
	suite.OutputDir = "./output"
}

func (suite *ZipTestSuite) TearDownTest() {
	//
}

func (suite *ZipTestSuite) TestUnzip() {
	paths, err := Unzip("testdata/test.zip", suite.OutputDir)
	suite.Nil(err)
	suite.Equal(paths, []string{"output/test.txt", "output/gophercolor16x16.png"})
}

func (suite *ZipTestSuite) TestIsDirectory() {
	isdir, err := IsDirectory("testdata")
	suite.Nil(err)
	suite.True(isdir)

	isdir, err = IsDirectory("testdata/hoge")
	suite.Nil(err)
	suite.False(isdir)

	isdir, err = IsDirectory("neverexists")
	suite.Error(err)
	suite.False(isdir)
}
