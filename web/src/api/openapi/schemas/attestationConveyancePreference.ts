/**
 * Generated by orval v6.15.0 🍺
 * Do not edit manually.
 * storyden
 * Storyden social API for building community driven platforms.
 * OpenAPI spec version: 1
 */

/**
 * https://www.w3.org/TR/webauthn-2/#enum-attestation-convey

 */
export type AttestationConveyancePreference =
  (typeof AttestationConveyancePreference)[keyof typeof AttestationConveyancePreference];

// eslint-disable-next-line @typescript-eslint/no-redeclare
export const AttestationConveyancePreference = {
  direct: "direct",
  enterprise: "enterprise",
  indirect: "indirect",
  none: "none",
} as const;
