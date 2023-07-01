import SeasonDialog from "@/components/Dialog/SeasonDialog";
import { render, screen, waitFor } from "@testing-library/react";
import { MovieType } from "@/types/movies";

describe("SeasonDialog", () => {
  test("fetches seasons and displays them in the dialog", async () => {
    const mockSelectedMovie: MovieType = {
      id: 1,
      title: "Mock Movie",
      type_code: "TV",
      description: "",
      release_date: "",
      runtime: 0
    };

    render(
      <SeasonDialog
        setNotifyState={jest.fn()}
        selectedMovie={mockSelectedMovie}
        setSelectedMovie={jest.fn()}
        open={true}
        setOpen={jest.fn()}
      />
    );

    await waitFor(() => {
      // Assert that the seasons are displayed
      expect(screen.getByText("Season 1")).toBeInTheDocument();
      expect(screen.getByText("Season 4")).toBeInTheDocument();

    });

    // Trigger an action within the dialog
    expect(screen.getByRole("link", { name: "Add Season" })).toHaveAttribute("href", "/admin/manage-catalogue/movies/seasons?movieId=1");

  })
});