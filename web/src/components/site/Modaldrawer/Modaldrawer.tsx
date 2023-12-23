import { PropsWithChildren } from "react";
import { Drawer } from "vaul";

import { CloseButton, Heading, UseDisclosureProps } from "src/theme/components";

import { Box, HStack, VStack } from "@/styled-system/jsx";

type Props = {
  title?: string;
} & UseDisclosureProps;

export function ModalDrawer({ children, ...props }: PropsWithChildren<Props>) {
  const onOpenChange = (open: boolean) => {
    if (open) props.onOpen?.();
    else props.onClose?.();
  };

  return (
    <>
      <Drawer.Root
        open={props.isOpen}
        onOpenChange={onOpenChange}
        // TODO: Scale background only on mobile.
        shouldScaleBackground={false}
      >
        <Drawer.Portal>
          <Drawer.Overlay className="modaldrawer__overlay" />
          <Drawer.Content className="modaldrawer__content">
            <VStack
              minHeight={{ base: "full", md: "0" }}
              minWidth={{ base: "full", md: "md" }}
              maxWidth={{ base: "full", md: "prose" }}
              borderTopRadius="xl"
              borderBottomRadius={{ base: "none", md: "xl" }}
              bgColor="gray.100"
              p="4"
            >
              <HStack w="full" justify="space-between">
                <Heading size="md">{props.title}</Heading>
                <CloseButton onClick={props.onClose} />
              </HStack>

              <Box h="full" w="full" pb="3">
                {children}
              </Box>
            </VStack>
          </Drawer.Content>
        </Drawer.Portal>
      </Drawer.Root>

      <style jsx global>{`
        .modaldrawer__overlay {
          position: fixed;
          inset: 0;
          background-color: var(--chakra-colors-blackAlpha-600);
          z-index: var(--z-index-overlay);
        }

        /* Modal mode - on desktop screens */
        @media screen and (min-width: 48em) {
          .modaldrawer__content {
            height: 100%;
            width: 100%;
            top: 0;
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            position: fixed;
            z-index: var(--z-index-modal);
          }
        }

        /* Drawer mode - on mobile screens */
        @media screen and (max-width: 48em) {
          .modaldrawer__content {
            height: 100%;
            width: 100%;
            top: 0;
            display: flex;
            flex-direction: column;
            position: fixed;
            margin-top: 3rem;
            max-height: 96%;
            z-index: var(--z-index-modal);
          }
        }
      `}</style>
    </>
  );
}
