import { render, screen, waitFor } from "@testing-library/react";
import Genres from "@/pages/genres";

describe("All Genres", () => {
  test("renders the all genres", async () => {
    render(<Genres />);
    await waitFor(() => {
      expect(screen.getByText("G1")).toBeInTheDocument(); // MOVIE

      expect(screen.getByText("G4")).toBeInTheDocument(); // TV

      expect(screen.getAllByRole("listitem")).toHaveLength(7)
    });
  });
});