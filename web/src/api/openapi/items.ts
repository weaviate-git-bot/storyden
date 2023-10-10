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
  InternalServerErrorResponse,
  ItemCreateBody,
  ItemCreateOKResponse,
  ItemGetOKResponse,
  ItemListOKResponse,
  ItemUpdateBody,
  ItemUpdateOKResponse,
  NotFoundResponse,
  UnauthorisedResponse,
} from "./schemas";

type AwaitedInput<T> = PromiseLike<T> | T;

type Awaited<O> = O extends AwaitedInput<infer T> ? T : never;

/**
 * Create a item to represent a piece of structured data such as an item in
a video game, an article of clothing, a product in a store, etc.

 */
export const itemCreate = (itemCreateBody: ItemCreateBody) => {
  return fetcher<ItemCreateOKResponse>({
    url: `/v1/items`,
    method: "post",
    headers: { "Content-Type": "application/json" },
    data: itemCreateBody,
  });
};

/**
 * List all items using the filtering options.
 */
export const itemList = () => {
  return fetcher<ItemListOKResponse>({ url: `/v1/items`, method: "get" });
};

export const getItemListKey = () => [`/v1/items`] as const;

export type ItemListQueryResult = NonNullable<
  Awaited<ReturnType<typeof itemList>>
>;
export type ItemListQueryError = NotFoundResponse | InternalServerErrorResponse;

export const useItemList = <
  TError = NotFoundResponse | InternalServerErrorResponse,
>(options?: {
  swr?: SWRConfiguration<Awaited<ReturnType<typeof itemList>>, TError> & {
    swrKey?: Key;
    enabled?: boolean;
  };
}) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false;
  const swrKey =
    swrOptions?.swrKey ?? (() => (isEnabled ? getItemListKey() : null));
  const swrFn = () => itemList();

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
 * Get a item by its URL slug.
 */
export const itemGet = (itemSlug: string) => {
  return fetcher<ItemGetOKResponse>({
    url: `/v1/items/${itemSlug}`,
    method: "get",
  });
};

export const getItemGetKey = (itemSlug: string) =>
  [`/v1/items/${itemSlug}`] as const;

export type ItemGetQueryResult = NonNullable<
  Awaited<ReturnType<typeof itemGet>>
>;
export type ItemGetQueryError =
  | UnauthorisedResponse
  | NotFoundResponse
  | InternalServerErrorResponse;

export const useItemGet = <
  TError =
    | UnauthorisedResponse
    | NotFoundResponse
    | InternalServerErrorResponse,
>(
  itemSlug: string,
  options?: {
    swr?: SWRConfiguration<Awaited<ReturnType<typeof itemGet>>, TError> & {
      swrKey?: Key;
      enabled?: boolean;
    };
  },
) => {
  const { swr: swrOptions } = options ?? {};

  const isEnabled = swrOptions?.enabled !== false && !!itemSlug;
  const swrKey =
    swrOptions?.swrKey ?? (() => (isEnabled ? getItemGetKey(itemSlug) : null));
  const swrFn = () => itemGet(itemSlug);

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
 * Update a item.
 */
export const itemUpdate = (
  itemSlug: string,
  itemUpdateBody: ItemUpdateBody,
) => {
  return fetcher<ItemUpdateOKResponse>({
    url: `/v1/items/${itemSlug}`,
    method: "patch",
    headers: { "Content-Type": "application/json" },
    data: itemUpdateBody,
  });
};
