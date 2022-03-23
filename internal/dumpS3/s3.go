package dumpS3

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/robfig/cron/v3"
	"gopkg.in/yaml.v2"
)

type S3Config struct {
	S3Endpoint  string `yaml:"s3Endpoint"`
	S3Region    string `yaml:"s3Region"`
	S3Bucket    string `yaml:"s3Bucket"`
	CronForSave string `yaml:"сronForSave"`
	FilePath    string `yaml:"filePath"`
}

// Config - contains all configuration parameters in Config package
type Config struct {
	CFG S3Config `yaml:"s3_config"`
}

func ReadConfigYML(filePath string) (cfg *Config, err error) {
	file, err := os.Open(filePath)
//	defer file.Close()
	if err != nil {
		return cfg, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}

func Dump() {
	s3Conf, err := ReadConfigYML("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	s3Token := os.Getenv("s3Token")
	s3Secret := os.Getenv("s3Secret")
	if s3Token == "" || s3Secret == "" {
		log.Fatal("provide token and secret for s3")
	}
	cron := cron.New()
	_, err = cron.AddFunc(s3Conf.CFG.CronForSave, func() {
		err := UploadToS3(
			s3Conf.CFG.S3Endpoint,
			s3Token,
			s3Secret,
			s3Conf.CFG.S3Region,
			s3Conf.CFG.S3Bucket,
			s3Conf.CFG.FilePath,
		)
		if err != nil {
			log.Println(err)
		}
	})
	if err != nil {
		log.Fatal(err)
	}
	cron.Start()
	log.Println("s3 dump inited")
}

func UploadToS3(endpoint, token, secret, region, bucket, filename string) error {
	fmt.Println("here!!!!!!!!!!!!!")
	awsSession, err := session.NewSession(
		aws.NewConfig().
			WithCredentials(
				credentials.NewStaticCredentials(
					token,
					secret,
					token)).
			WithEndpoint(endpoint).
			WithDisableSSL(false).
			WithS3ForcePathStyle(true).
			WithRegion(region),
	)
	if err != nil {
		return fmt.Errorf("uploadToS3 %s", err.Error())
	}

	s3Client := s3.New(awsSession, aws.NewConfig().WithEndpoint(endpoint))

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Body:   file,
		Bucket: &bucket,
		Key:    &filename,
	})

	if err != nil {
		return fmt.Errorf("uploadToS3 %s", err.Error())
	}
	return nil
}
