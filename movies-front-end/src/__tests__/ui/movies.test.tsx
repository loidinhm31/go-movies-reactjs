import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import Movies from "@/pages/movies";
import userEvent from "@testing-library/user-event";
import { TabMovie } from "@/components/Tab/TabMovie";
import { TabTvSeries } from "@/components/Tab/TabTvSeries";
import { BuyCollection } from "@/components/Payment/BuyCollection";
import { MovieType } from "@/types/movies";
import WatchMovie from "@/components/Movie/WatchMovie";
import { useRouter } from "next/router";
import { useHasUsername } from "@/hooks/auth/useHasUsername";
import { Views } from "@/components/Views";

describe("all movies", () => {
  it("changes the active tab when clicked", async () => {
    render(<Movies />);

    await waitFor(() => {
      const titleElement = screen.getByText("All Movies");

      const moviesTab = screen.getByRole("tab", { name: /movies/i });
      const tvSeriesTab = screen.getByRole("tab", { name: /tv series/i });

      // Initially, the movies tab should be active
      expect(titleElement).toBeInTheDocument();

      expect(moviesTab).toBeInTheDocument();
      expect(tvSeriesTab).toBeInTheDocument();

      expect(moviesTab).toHaveAttribute("aria-selected", "true");
      expect(tvSeriesTab).toHaveAttribute("aria-selected", "false");

      // Click on the TV series tab
      fireEvent.click(tvSeriesTab);

      // After clicking, the TV series tab should be active
      expect(moviesTab).toHaveAttribute("aria-selected", "false");
      expect(tvSeriesTab).toHaveAttribute("aria-selected", "true");
    });
  });

  it("renders Movies component and switches tabs correctly", async () => {
    render(<Movies />);

    await waitFor(() => {
      // Verify that the initial tab is selected correctly
      const moviesTab = screen.getByRole("tab", { name: "Movies" });
      const tvSeriesTab = screen.getByRole("tab", { name: "TV Series" });
      expect(moviesTab).toHaveAttribute("aria-selected", "true");
      expect(tvSeriesTab).toHaveAttribute("aria-selected", "false");

      // Simulate clicking on the TV Series tab
      userEvent.click(tvSeriesTab);

      // Verify that the tab value is updated correctly
      expect(moviesTab).toHaveAttribute("aria-selected", "false");
      expect(tvSeriesTab).toHaveAttribute("aria-selected", "true");

      // Verify that the TabMovie component is rendered when the Movies tab is selected
      const tabMovieComponent = screen.getByTestId("TheatersIcon");
      expect(tabMovieComponent).toBeInTheDocument();
    });
  });

  it("tab movie component", async () => {
    render(<TabMovie />);

    await waitFor(() => {
      const keyword = screen.getByRole("textbox", { name: "Keyword" });

      const movieName = screen.getByText("Test movie 1");

      expect(keyword).toBeInTheDocument();

      expect(movieName).toBeInTheDocument();
    });
  });

  it("tab tv component", async () => {
    render(<TabTvSeries />);

    await waitFor(() => {
      const keyword = screen.getByRole("textbox", { name: "Keyword" });

      const movieName = screen.getByText("Test tv 3");

      expect(keyword).toBeInTheDocument();

      expect(movieName).toBeInTheDocument();
    });
  });
});


jest.mock("next/router", () => ({
  ...(jest.requireActual("next/router") as object),
  useRouter: jest.fn()
}));

jest.mock("src/hooks/auth/useHasUsername", () => ({
  useHasUsername: jest.fn()
}));

describe("Movie Component", () => {

  test("renders the movie component with the provided author and movieId", async () => {

    (useRouter as jest.Mock).mockImplementation(() => ({
      query: { id: "1" }
    }));

    const mockAuthor = "JohnDoe";
    const mockUseHasUsername = jest.fn(() => mockAuthor);
    (useHasUsername as jest.Mock).mockReturnValue(mockAuthor);

    render(<WatchMovie author={mockAuthor} movieId={1} />);

    await waitFor(() => {
      const date = screen.getByText("June 5th, 2014 |");
      const description = screen.getByText("Movie - Test movie 1");

      expect(date).toBeInTheDocument();
      expect(description).toBeInTheDocument();

    });
  });

  test("renders view component", async () => {

    render(<Views movieId={1} wasMutateView={true} setWasMuateView={jest.fn} />);

    await waitFor(() => {
      const views = screen.getByText("5 views");

      expect(views).toBeInTheDocument();

    });
  });
});

describe("Buy Collection Component", () => {

  test("renders correct button when not paid and not collected", async () => {
    const movie: MovieType = {
      id: 2, title: "test", description: "test desc", type_code: "MOVIE", price: 9.99, release_date: "", runtime: 0
    };

    render(<BuyCollection movie={movie} />);

    await waitFor(() => {
      // Expect "Price" to be rendered
      expect(screen.getByText("9.99")).toBeInTheDocument();

      // Expect "Buy to Watch" button to be rendered
      expect(screen.getByRole("button", { name: "Buy to Watch" })).toBeInTheDocument();
    });
  });

  test("renders correct button when paid and collected", async () => {
    const movie: MovieType = {
      id: 1, title: "test", description: "test desc", type_code: "MOVIE", price: 9.99, release_date: "", runtime: 0
    };

    render(<BuyCollection movie={movie} />);

    // Wait for the component to update after checking payment and collection
    await waitFor(() => {
      expect(screen.getByRole("button", { name: "Collected" })).toBeInTheDocument();
    });
  });

  test("clicking Add to Collection triggers addCollection", async () => {
    const movie: MovieType = {
      id: 2, title: "test", description: "test desc", type_code: "MOVIE", release_date: "", runtime: 0
    };
    render(<BuyCollection movie={movie} />);

    await waitFor(async () => {
      userEvent.click(screen.getByRole("button", { name: "Add to Collection" }));

      await waitFor(() => {
        // Expect the collection to be added
        expect(screen.getByRole("button", { name: "Collected" })).toBeInTheDocument();
      });
    });
  });

  test("clicking Collected triggers deleteCollection", async () => {
    const movie: MovieType = {
      id: 1, title: "test", description: "test desc", type_code: "MOVIE", release_date: "", runtime: 0
    };
    render(<BuyCollection movie={movie} />);

    // Wait for the deleteCollection API request to be sent and processed
    await waitFor(async () => {
      userEvent.click(screen.getByRole("button", { name: "Collected" }));

      await waitFor(() => {
        // Expect the collection to be removed
        expect(screen.getByRole("button", { name: "Add to Collection" })).toBeInTheDocument();
      });
    });
  });
});