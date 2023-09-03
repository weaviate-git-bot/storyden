/**
 * Generated by orval v6.15.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import useSwr from "swr";
import type { Key, SWRConfiguration } from "swr";

import { fetcher } from "../client";

import type { GetInfoOKResponse, InternalServerErrorResponse } from "./schemas";

/**
 * The version number includes the date and time of the release build as
well as a short representation of the Git commit hash.

 * @summary Get the software version string.
 */
export const getVersion = () => {
  return fetcher<string>({ url: `/version`, method: "get" });
};

export const getGetVersionKey = () => [`/version`] as const;

export type GetVersionQueryResult = NonNullable<
  Awaited<ReturnType<typeof getVersion>>
>;
export type GetVersionQueryError = unknown;

export const useGetVersion = <TError = unknown>(options?: {
  swr?: SWRConfiguration<Awaited<ReturnType<typeof getVersion>>, TError> & {
    swrKey?: Key;
    enabled?: boolean;
  };
}) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ?? (() => (isEnabled ? getGetVersionKey() : null));
  const swrFn = () => getVersion();

  const query = useSwr<Awaited<ReturnType<typeof swrFn>>, TError>(
    swrKey,
    swrFn,
    swrOptions,
  );

  return {
    swrKey,
    ...query,
  };
};

/**
 * Note: the generator creates a `map[string]interface{}` if this is set to
`application/json`... so I'm just using plain text for now.

 * @summary Get the OpenAPI 3.0 specification as JSON.
 */
export const getSpec = () => {
  return fetcher<string>({ url: `/openapi.json`, method: "get" });
};

export const getGetSpecKey = () => [`/openapi.json`] as const;

export type GetSpecQueryResult = NonNullable<
  Awaited<ReturnType<typeof getSpec>>
>;
export type GetSpecQueryError = unknown;

export const useGetSpec = <TError = unknown>(options?: {
  swr?: SWRConfiguration<Awaited<ReturnType<typeof getSpec>>, TError> & {
    swrKey?: Key;
    enabled?: boolean;
  };
}) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ?? (() => (isEnabled ? getGetSpecKey() : null));
  const swrFn = () => getSpec();

  const query = useSwr<Awaited<ReturnType<typeof swrFn>>, TError>(
    swrKey,
    swrFn,
    swrOptions,
  );

  return {
    swrKey,
    ...query,
  };
};

/**
 * Get the basic forum installation info such as title, description, etc.

 */
export const getInfo = () => {
  return fetcher<GetInfoOKResponse>({ url: `/v1/info`, method: "get" });
};

export const getGetInfoKey = () => [`/v1/info`] as const;

export type GetInfoQueryResult = NonNullable<
  Awaited<ReturnType<typeof getInfo>>
>;
export type GetInfoQueryError = InternalServerErrorResponse;

export const useGetInfo = <TError = InternalServerErrorResponse>(options?: {
  swr?: SWRConfiguration<Awaited<ReturnType<typeof getInfo>>, TError> & {
    swrKey?: Key;
    enabled?: boolean;
  };
}) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ?? (() => (isEnabled ? getGetInfoKey() : null));
  const swrFn = () => getInfo();

  const query = useSwr<Awaited<ReturnType<typeof swrFn>>, TError>(
    swrKey,
    swrFn,
    swrOptions,
  );

  return {
    swrKey,
    ...query,
  };
};
