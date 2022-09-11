/**
 * Generated by orval v6.9.6 🍺
 * Do not edit manually.
 * storyden
 * OpenAPI spec version: 1
 */
import type { CommonProperties } from "./commonProperties";
import type { ThreadReferenceAllOf } from "./threadReferenceAllOf";

/**
 * A thread reference includes most of the information about a thread but
does not include the posts within the thread. Useful for rendering large
lists of threads or other situations when you don't need the full data.

 */
export type ThreadReference = CommonProperties & ThreadReferenceAllOf;
