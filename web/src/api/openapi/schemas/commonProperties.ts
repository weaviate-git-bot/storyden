/**
 * Generated by orval v6.9.6 🍺
 * Do not edit manually.
 * storyden
 * OpenAPI spec version: 1
 */
import type { Identifier } from "./identifier";

export interface CommonProperties {
  id: Identifier;
  /** The time the resource was created. */
  createdAt: string;
  /** The time the resource was updated. */
  updatedAt: string;
  /** The time the resource was soft-deleted. */
  deletedAt?: string;
}
