import { TimeblockType } from "./components/Timeblock";

export enum Event {
  NewBlock = 'newblock',

};

/**
 * Wrapper function to create and dispatch an event
 *
 * @param name Name of the event
 * @param payload Payload within the event
 */
export function fireEvent(name: Event, payload: null | TimeblockType) {
  const event = new CustomEvent(name, { detail: payload });
  window.dispatchEvent(event);
}