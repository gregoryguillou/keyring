import { ProviderMessage } from "./provider";
import { AccountMessage } from "./account";
import { KeyringMessage } from "./keyring";
import { ConfigMessage } from "./configuration";

export const uuid = "589c80c1eb85413d";
export const defaultTimeoutMilliseconds = 30000;

export type MessageType =
  | ProviderMessage
  | KeyringMessage
  | ConfigMessage
  | AccountMessage;

export type WindowMessageType = MessageType & {
  uuid: string;
};

export const sendMessage = (msg: MessageType): void => {
  const keyring = document
    .querySelector("#starknetburner")
    ?.querySelector<HTMLIFrameElement>("#iframe");
  return keyring?.contentWindow?.postMessage({ ...msg, uuid }, "*");
};

export const waitForMessage = async <
  K extends MessageType["type"],
  T extends { type: K } & MessageType
>(
  type: K,
  predicate: (x: T) => boolean = () => true
): Promise<T extends { data: infer S } ? S : undefined> => {
  return new Promise((resolve, reject) => {
    const pid = setTimeout(
      () => reject(new Error("Timeout")),
      defaultTimeoutMilliseconds
    );
    const handler = (event: MessageEvent<WindowMessageType>) => {
      if (
        event.data.type === type &&
        event.data.uuid === uuid &&
        predicate(event.data as any)
      ) {
        clearTimeout(pid);
        window.removeEventListener("message", handler);
        return resolve(
          ("data" in event.data ? event.data.data : undefined) as any
        );
      }
    };
    window.addEventListener("message", handler);
  });
};
