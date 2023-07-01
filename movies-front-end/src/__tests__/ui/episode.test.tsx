import { render, screen, waitFor } from "@testing-library/react";
import WatchEpisode from "@/components/Movie/WatchEpisode";
import { useRouter } from "next/router";
import { useHasUsername } from "@/hooks/auth/useHasUsername";


jest.mock("next/router", () => ({
  ...(jest.requireActual("next/router") as object),
  useRouter: jest.fn()
}));

jest.mock("src/hooks/auth/useHasUsername", () => ({
  useHasUsername: jest.fn()
}));

describe("Watch Episode", () => {
  test("renders the component with correct data", async () => {
    (useRouter as jest.Mock).mockImplementation(() => ({
      query: { id: "5" }
    }));

    const mockAuthor = "JohnDoe";
    const mockUseHasUsername = jest.fn(() => mockAuthor);
    (useHasUsername as jest.Mock).mockReturnValue(mockAuthor);

    render(<WatchEpisode author={mockAuthor} movieId={5} episodeId={1} />);

    await waitFor(() => {
      expect(screen.getByText("TV Series - Test tv 5")).toBeInTheDocument();

      // Expect the episode name to be rendered
      expect(screen.getByText("Season 1 - E1: Test episode 1")).toBeInTheDocument();
    });
  });
});