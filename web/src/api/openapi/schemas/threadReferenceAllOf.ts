/**
 * Generated by orval v6.9.6 🍺
 * Do not edit manually.
 * storyden
 * OpenAPI spec version: 1
 */
import type { ProfileReference } from "./profileReference";
import type { Category } from "./category";
import type { React } from "./react";

export type ThreadReferenceAllOf = {
  /** The title of the thread. */
  title: string;
  /** A URL friendly slug which is prepended with the post ID
for uniqueness and sortability.
 */
  readonly slug: string;
  /** A short version of the thread's body text for use in previews.
   */
  readonly short: string;
  /** Whether the thread is pinned in this category. */
  pinned: boolean;
  author: ProfileReference;
  /** A list of tags associated with the thread. */
  tags: string[];
  /** The number of posts under this thread. */
  readonly posts: number;
  category: Category;
  /** A list of reactions this post has had from people. */
  reacts: React[];
};
