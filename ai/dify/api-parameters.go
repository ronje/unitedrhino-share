package dify

import (
	"encoding/json"
	"fmt"
)

type T struct {
	OpeningStatement              string        `json:"opening_statement"`
	SuggestedQuestions            []interface{} `json:"suggested_questions"`
	SuggestedQuestionsAfterAnswer struct {
		Enabled bool `json:"enabled"`
	} `json:"suggested_questions_after_answer"`
	SpeechToText struct {
		Enabled bool `json:"enabled"`
	} `json:"speech_to_text"`
	TextToSpeech struct {
		Enabled  bool   `json:"enabled"`
		Voice    string `json:"voice"`
		Language string `json:"language"`
	} `json:"text_to_speech"`
	RetrieverResource struct {
		Enabled bool `json:"enabled"`
	} `json:"retriever_resource"`
	AnnotationReply struct {
		Enabled bool `json:"enabled"`
	} `json:"annotation_reply"`
	MoreLikeThis struct {
		Enabled bool `json:"enabled"`
	} `json:"more_like_this"`
	UserInputForm          []interface{} `json:"user_input_form"`
	SensitiveWordAvoidance struct {
		Enabled bool          `json:"enabled"`
		Type    string        `json:"type"`
		Configs []interface{} `json:"configs"`
	} `json:"sensitive_word_avoidance"`
	FileUpload struct {
		Image struct {
			Detail          string   `json:"detail"`
			Enabled         bool     `json:"enabled"`
			NumberLimits    int      `json:"number_limits"`
			TransferMethods []string `json:"transfer_methods"`
		} `json:"image"`
		Enabled                  bool          `json:"enabled"`
		AllowedFileTypes         []interface{} `json:"allowed_file_types"`
		AllowedFileExtensions    []string      `json:"allowed_file_extensions"`
		AllowedFileUploadMethods []string      `json:"allowed_file_upload_methods"`
		NumberLimits             int           `json:"number_limits"`
	} `json:"file_upload"`
	SystemParameters struct {
		ImageFileSizeLimit      int `json:"image_file_size_limit"`
		VideoFileSizeLimit      int `json:"video_file_size_limit"`
		AudioFileSizeLimit      int `json:"audio_file_size_limit"`
		FileSizeLimit           int `json:"file_size_limit"`
		WorkflowFileUploadLimit int `json:"workflow_file_upload_limit"`
	} `json:"system_parameters"`
}

type GetParametersResponse struct {
	OpeningStatement              string `json:"opening_statement"`
	SuggestedQuestions            []any  `json:"suggested_questions"`
	SuggestedQuestionsAfterAnswer struct {
		Enabled bool `json:"enabled"`
	} `json:"suggested_questions_after_answer"`
	SpeechToText struct {
		Enabled bool `json:"enabled"`
	} `json:"speech_to_text"`
	TextToSpeech struct {
		Enabled  bool   `json:"enabled"`
		Voice    string `json:"voice"`
		Language string `json:"language"`
	} `json:"text_to_speech"`
	RetrieverResource struct {
		Enabled bool `json:"enabled"`
	} `json:"retriever_resource"`
	AnnotationReply struct {
		Enabled bool `json:"enabled"`
	} `json:"annotation_reply"`
	MoreLikeThis struct {
		Enabled bool `json:"enabled"`
	} `json:"more_like_this"`
	UserInputForm []struct {
		Paragraph struct {
			Label    string `json:"label"`
			Variable string `json:"variable"`
			Required bool   `json:"required"`
			Default  string `json:"default"`
		} `json:"paragraph"`
	} `json:"user_input_form"`
	SensitiveWordAvoidance struct {
		Enabled bool   `json:"enabled"`
		Type    string `json:"type"`
		Configs []any  `json:"configs"`
	} `json:"sensitive_word_avoidance"`
	FileUpload struct {
		Image struct {
			Detail          string   `json:"detail"`
			Enabled         bool     `json:"enabled"`
			NumberLimits    int      `json:"number_limits"`
			TransferMethods []string `json:"transfer_methods"`
		} `json:"image"`
		Enabled                  bool          `json:"enabled"`
		AllowedFileTypes         []interface{} `json:"allowed_file_types"`
		AllowedFileExtensions    []string      `json:"allowed_file_extensions"`
		AllowedFileUploadMethods []string      `json:"allowed_file_upload_methods"`
		NumberLimits             int           `json:"number_limits"`
	} `json:"file_upload"`
	SystemParameters struct {
		ImageFileSizeLimit      int `json:"image_file_size_limit"`
		VideoFileSizeLimit      int `json:"video_file_size_limit"`
		AudioFileSizeLimit      int `json:"audio_file_size_limit"`
		FileSizeLimit           int `json:"file_size_limit"`
		WorkflowFileUploadLimit int `json:"workflow_file_upload_limit"`
	} `json:"system_parameters"`
}

func (dc *DifyClient) GetParameters() (result GetParametersResponse, err error) {
	api := dc.GetAPI(API_PARAMETERS)
	code, body, err := SendGetRequestToAPI(dc, api)

	err = CommonRiskForSendRequest(code, err)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal the response: %v", err)
	}
	return result, nil
}
