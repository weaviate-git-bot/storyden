/**
 * Generated by orval v6.9.6 🍺
 * Do not edit manually.
 * storyden
 * OpenAPI spec version: 1
 */
import useSwr from "swr";
import type { SWRConfiguration, Key } from "swr";
import type {
  AuthSuccessResponse,
  BadRequestResponse,
  InternalServerErrorResponse,
  WebAuthnMakeCredentialBody,
} from "./schemas";
import { fetcher } from "../client";

/**
 * @summary Complete WebAuthn registration by creating a new credential.
 */
export const webAuthnMakeCredential = (
  webAuthnMakeCredentialBody: WebAuthnMakeCredentialBody
) => {
  return fetcher<AuthSuccessResponse>({
    url: `/v1/auth/webauthn/make`,
    method: "get",
    headers: { "Content-Type": "application/json" },
  });
};

export const getWebAuthnMakeCredentialKey = (
  webAuthnMakeCredentialBody: WebAuthnMakeCredentialBody
) => [`/v1/auth/webauthn/make`, webAuthnMakeCredentialBody];

export type WebAuthnMakeCredentialQueryResult = NonNullable<
  Awaited<ReturnType<typeof webAuthnMakeCredential>>
>;
export type WebAuthnMakeCredentialQueryError =
  | BadRequestResponse
  | InternalServerErrorResponse;

export const useWebAuthnMakeCredential = <
  TError = BadRequestResponse | InternalServerErrorResponse
>(
  webAuthnMakeCredentialBody: WebAuthnMakeCredentialBody,
  options?: {
    swr?: SWRConfiguration<
      Awaited<ReturnType<typeof webAuthnMakeCredential>>,
      TError
    > & { swrKey?: Key; enabled?: boolean };
  }
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ??
    (() =>
      isEnabled
        ? getWebAuthnMakeCredentialKey(webAuthnMakeCredentialBody)
        : null);
  const swrFn = () => webAuthnMakeCredential(webAuthnMakeCredentialBody);

  const query = useSwr<Awaited<ReturnType<typeof swrFn>>, TError>(
    swrKey,
    swrFn,
    swrOptions
  );

  return {
    swrKey,
    ...query,
  };
};
