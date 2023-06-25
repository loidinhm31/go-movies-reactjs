// jest.setup.js
import "@testing-library/jest-dom/extend-expect";

// polyfill for node
import { TextDecoderStream } from "@stardazed/streams-text-encoding";
import { TextDecoder, TextEncoder } from "util";
import { ReadableStream } from "web-streams-polyfill/es6";
import ResizeObserver from 'resize-observer-polyfill';
import "jest-canvas-mock";
import { server } from "@/__tests__/__mocks__/msw/server";

global.TextEncoder = TextEncoder;
global.TextDecoderStream = TextDecoderStream;
global.TextDecoder = TextDecoder;
global.ReadableStream = ReadableStream;
global.ResizeObserver = ResizeObserver;

// Establish API mocking before all tests.
beforeAll(() => server.listen());

// Reset any request handlers that we may add during the tests,
// so they don't affect other tests.
afterEach(() => server.resetHandlers());

// Clean up after the tests are finished.
afterAll(() => server.close());