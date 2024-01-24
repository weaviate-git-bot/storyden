/**
 * Generated by orval v6.17.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */
import useSwr from "swr";
import type { Key, SWRConfiguration } from "swr";

import { fetcher } from "../client";

import type {
  ClusterAddChildOKResponse,
  ClusterAddItemOKResponse,
  ClusterCreateBody,
  ClusterCreateOKResponse,
  ClusterDeleteOKResponse,
  ClusterDeleteParams,
  ClusterGetOKResponse,
  ClusterListOKResponse,
  ClusterListParams,
  ClusterRemoveChildOKResponse,
  ClusterRemoveItemOKResponse,
  ClusterUpdateBody,
  ClusterUpdateOKResponse,
  InternalServerErrorResponse,
  NotFoundResponse,
  UnauthorisedResponse,
} from "./schemas";

type AwaitedInput<T> = PromiseLike<T> | T;

type Awaited<O> = O extends AwaitedInput<infer T> ? T : never;

/**
 * Create a cluster for grouping items and other clusters together.

 */
export const clusterCreate = (clusterCreateBody: ClusterCreateBody) => {
  return fetcher<ClusterCreateOKResponse>({
    url: `/v1/clusters`,
    method: "post",
    headers: { "Content-Type": "application/json" },
    data: clusterCreateBody,
  });
};

/**
 * List all clusters.
 */
export const clusterList = (params?: ClusterListParams) => {
  return fetcher<ClusterListOKResponse>({
    url: `/v1/clusters`,
    method: "get",
    params,
  });
};

export const getClusterListKey = (params?: ClusterListParams) =>
  [`/v1/clusters`, ...(params ? [params] : [])] as const;

export type ClusterListQueryResult = NonNullable<
  Awaited<ReturnType<typeof clusterList>>
>;
export type ClusterListQueryError =
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useClusterList = <
  TError = NotFoundResponse | InternalServerErrorResponse,
>(
  params?: ClusterListParams,
  options?: {
    swr?: SWRConfiguration<Awaited<ReturnType<typeof clusterList>>, TError> & {
      swrKey?: Key;
      enabled?: boolean;
    };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ??
    (() => (isEnabled ? getClusterListKey(params) : null));
  const swrFn = () => clusterList(params);

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
 * Get a cluster by its URL slug.
 */
export const clusterGet = (clusterSlug: string) => {
  return fetcher<ClusterGetOKResponse>({
    url: `/v1/clusters/${clusterSlug}`,
    method: "get",
  });
};

export const getClusterGetKey = (clusterSlug: string) =>
  [`/v1/clusters/${clusterSlug}`] as const;

export type ClusterGetQueryResult = NonNullable<
  Awaited<ReturnType<typeof clusterGet>>
>;
export type ClusterGetQueryError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useClusterGet = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  clusterSlug: string,
  options?: {
    swr?: SWRConfiguration<Awaited<ReturnType<typeof clusterGet>>, TError> & {
      swrKey?: Key;
      enabled?: boolean;
    };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false && !!clusterSlug;
  const swrKey =
    swrOptions?.swrKey ??
    (() => (isEnabled ? getClusterGetKey(clusterSlug) : null));
  const swrFn = () => clusterGet(clusterSlug);

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
 * Update a cluster.
 */
export const clusterUpdate = (
  clusterSlug: string,
  clusterUpdateBody: ClusterUpdateBody,
) => {
  return fetcher<ClusterUpdateOKResponse>({
    url: `/v1/clusters/${clusterSlug}`,
    method: "patch",
    headers: { "Content-Type": "application/json" },
    data: clusterUpdateBody,
  });
};

/**
 * Delete a cluster and move all children to its parent or root.
 */
export const clusterDelete = (
  clusterSlug: string,
  params?: ClusterDeleteParams,
) => {
  return fetcher<ClusterDeleteOKResponse>({
    url: `/v1/clusters/${clusterSlug}`,
    method: "delete",
    params,
  });
};

/**
 * Add an asset to a cluster.
 */
export const clusterAddAsset = (clusterSlug: string, id: string) => {
  return fetcher<ClusterUpdateOKResponse>({
    url: `/v1/clusters/${clusterSlug}/assets/${id}`,
    method: "put",
  });
};

/**
 * Remove an asset from a cluster.
 */
export const clusterRemoveAsset = (clusterSlug: string, id: string) => {
  return fetcher<ClusterUpdateOKResponse>({
    url: `/v1/clusters/${clusterSlug}/assets/${id}`,
    method: "delete",
  });
};

/**
 * Set a cluster's parent to the specified cluster
 */
export const clusterAddCluster = (
  clusterSlug: string,
  clusterSlugChild: string,
) => {
  return fetcher<ClusterAddChildOKResponse>({
    url: `/v1/clusters/${clusterSlug}/clusters/${clusterSlugChild}`,
    method: "put",
  });
};

/**
 * Remove a cluster from its parent cluster and back to the top level.

 */
export const clusterRemoveCluster = (
  clusterSlug: string,
  clusterSlugChild: string,
) => {
  return fetcher<ClusterRemoveChildOKResponse>({
    url: `/v1/clusters/${clusterSlug}/clusters/${clusterSlugChild}`,
    method: "delete",
  });
};

/**
 * Add an item to a cluster.
 */
export const clusterAddItem = (clusterSlug: string, itemSlug: string) => {
  return fetcher<ClusterAddItemOKResponse>({
    url: `/v1/clusters/${clusterSlug}/items/${itemSlug}`,
    method: "put",
  });
};

/**
 * Remove an item from a cluster.
 */
export const clusterRemoveItem = (clusterSlug: string, itemSlug: string) => {
  return fetcher<ClusterRemoveItemOKResponse>({
    url: `/v1/clusters/${clusterSlug}/items/${itemSlug}`,
    method: "delete",
  });
};
