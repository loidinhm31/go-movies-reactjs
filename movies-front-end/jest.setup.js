// jest.setup.js
import "@testing-library/jest-dom/extend-expect";

// polyfill for node
import { TextDecoderStream } from "@stardazed/streams-text-encoding";
import { TextDecoder, TextEncoder } from "util";
import { ReadableStream } from "web-streams-polyfill/es6";

global.TextEncoder = TextEncoder;
global.TextDecoderStream = TextDecoderStream;
global.TextDecoder = TextDecoder;
global.ReadableStream = ReadableStream;

const CONSOLE_FAIL_TYPES = ["error", "warn"];

// Throw errors when a `console.error` or `console.warn` happens
// by overriding the functions.
// If the warning/error is intentional, then catch it and expect for it, like:
//
//  jest.spyOn(console, 'warn').mockImplementation();
//  ...
//  expect(console.warn).toHaveBeenCalledWith(expect.stringContaining('Empty titles are deprecated.'));
CONSOLE_FAIL_TYPES.forEach((type) => {
  const orig_f = console[type];
  console[type] = (...message) => {
    orig_f(...message);
    throw new Error(`Failing due to console.${type} while running test!\n\n${message.join(" ")}`);
  };
});