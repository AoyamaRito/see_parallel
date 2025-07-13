package analyzer

import (
	"context"
	"fmt"

	"see_parallel/internal/config"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiClient struct {
	client *genai.Client
	ctx    context.Context
}

func NewGeminiClient(ctx context.Context) (*GeminiClient, error) {
	apiKey := config.GetAPIKey()
	if apiKey == "" {
		return nil, fmt.Errorf("APIキーが設定されていません。'see_parallel api set \"YOUR_API_KEY\"' または環境変数 GEMINI_API_KEY を設定してください")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %v", err)
	}

	return &GeminiClient{
		client: client,
		ctx:    ctx,
	}, nil
}

func (g *GeminiClient) Analyze(question, content string, deep bool) (string, error) {
	var model *genai.GenerativeModel
	if deep {
		model = g.client.GenerativeModel("gemini-2.0-flash-exp")
	} else {
		model = g.client.GenerativeModel("gemini-1.5-flash")
	}

	contextInfo := config.GetContext()
	var prompt string
	if contextInfo != "" {
		prompt = fmt.Sprintf("【文脈情報】\n%s\n\n【質問】\n%s\n\n【コンテンツ】\n%s", contextInfo, question, content)
	} else {
		prompt = fmt.Sprintf("%s\n\n%s", question, content)
	}
	
	resp, err := model.GenerateContent(g.ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %v", err)
	}

	if len(resp.Candidates) == 0 {
		return "", fmt.Errorf("no response from Gemini")
	}

	var result string
	for _, part := range resp.Candidates[0].Content.Parts {
		if text, ok := part.(genai.Text); ok {
			result += string(text)
		}
	}

	return result, nil
}

func (g *GeminiClient) Close() error {
	return g.client.Close()
}