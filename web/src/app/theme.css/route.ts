import { flatten, zip } from "lodash";
import { NextResponse } from "next/server";

import { getColourVariants } from "src/utils/colour";
import { getInfo } from "src/utils/info";

export const dynamic = "force-dynamic";

/**
 *
 * @returns A fully static server-side-rendered CSS document that uses the
 *          Storyden installation's accent colour set by the administrator.
 */

export async function GET() {
  const info = await getInfo();

  const cv = getColourVariants(info.accent_colour);

  const rules = Object.entries(cv).map(([k, v]) => `${k}: ${v};`);

  const document = css`
    :root {
      ${rules.join("\n      ")}
    }
  `;

  return new NextResponse(document, {
    headers: {
      "Content-Type": "text/css",
    },
  });
}

// NOTE: literally just so we get syntax highlighting above...
function css(s: TemplateStringsArray, ...v: string[]): string {
  return flatten(zip(s, v)).join("");
}
