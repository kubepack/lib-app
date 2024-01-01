// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves an authorization token. An authorization token represents your IAM
// authentication credentials and can be used to access any Amazon ECR registry
// that your IAM principal has access to. The authorization token is valid for 12
// hours. The authorizationToken returned is a base64 encoded string that can be
// decoded and used in a docker login command to authenticate to a registry. The
// CLI offers an get-login-password command that simplifies the login process. For
// more information, see Registry authentication (https://docs.aws.amazon.com/AmazonECR/latest/userguide/Registries.html#registry_auth)
// in the Amazon Elastic Container Registry User Guide.
func (c *Client) GetAuthorizationToken(ctx context.Context, params *GetAuthorizationTokenInput, optFns ...func(*Options)) (*GetAuthorizationTokenOutput, error) {
	if params == nil {
		params = &GetAuthorizationTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAuthorizationToken", params, optFns, c.addOperationGetAuthorizationTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAuthorizationTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAuthorizationTokenInput struct {

	// A list of Amazon Web Services account IDs that are associated with the
	// registries for which to get AuthorizationData objects. If you do not specify a
	// registry, the default registry is assumed.
	//
	// Deprecated: This field is deprecated. The returned authorization token can be
	// used to access any Amazon ECR registry that the IAM principal has access to,
	// specifying a registry ID doesn't change the permissions scope of the
	// authorization token.
	RegistryIds []string

	noSmithyDocumentSerde
}

type GetAuthorizationTokenOutput struct {

	// A list of authorization token data objects that correspond to the registryIds
	// values in the request.
	AuthorizationData []types.AuthorizationData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAuthorizationTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAuthorizationToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAuthorizationToken{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAuthorizationToken"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAuthorizationToken(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetAuthorizationToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAuthorizationToken",
	}
}
