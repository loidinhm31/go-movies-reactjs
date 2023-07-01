import { CollectionMovieTab } from "@/components/Tab/CollectionMovieTab";
import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { CollectionEpisodeTab } from "@/components/Tab/CollectionEpisodeTab";

describe("CollectionMovieTab", () => {
  test("renders movie cards correctly", async () => {
    render(
      <CollectionMovieTab />
    );

    // Wait for the API response to be resolved and movie cards to be rendered
    await screen.findAllByRole("link");

    // Verify that movie cards are rendered with the correct data
    expect(screen.getByText("Movie 1")).toBeInTheDocument();
    expect(screen.getByText("9.99 USD")).toBeInTheDocument();
    expect(screen.getByText("Movie 1 description")).toBeInTheDocument();

    expect(screen.getByText("Movie 2")).toBeInTheDocument();
    expect(screen.getByText("FREE")).toBeInTheDocument();
    expect(screen.getByText("Movie 2 description")).toBeInTheDocument();
  });

  test("changes page size correctly", async () => {
    render(
      <CollectionMovieTab />
    );

    await waitFor(() => {
      expect(screen.getByRole("button", { name: "9" })).toBeInTheDocument();

    });

    userEvent.click(screen.getByRole("button", { name: "9" }));
    userEvent.click(screen.getByRole("option", { name: "18" }));

    await waitFor(() => {
      expect(screen.getByRole("button", { name: "18" })).toBeInTheDocument();
    })
  });
});

describe("CollectionEpisodeTab", () => {
  test("should render the episodes correctly", async () => {
    // Render the component wrapped in SWRConfig to provide the SWR context
    render(
      <CollectionEpisodeTab />
    );

    // Wait for the API response and the component to update
    await screen.findAllByRole("img"); // Wait for images to load
    await screen.findByText("Episode 1"); // Wait for episode title to appear

    // Check if the episodes are rendered correctly
    expect(screen.getByText("Episode 1")).toBeInTheDocument();
    expect(screen.getByText("Episode 2")).toBeInTheDocument();
  });

  it("should change the page size when the select input is changed", async () => {
    render(
      <CollectionEpisodeTab />

    );

    // Wait for the API response and the component to update
    await screen.findAllByRole("img"); // Wait for images to load

    await waitFor(() => {
      expect(screen.getByRole("button", { name: "9" })).toBeInTheDocument();

    });

    userEvent.click(screen.getByRole("button", { name: "9" }));
    userEvent.click(screen.getByRole("option", { name: "18" }));

    await waitFor(() => {
      expect(screen.getByRole("button", { name: "18" })).toBeInTheDocument();
    })
  });
});