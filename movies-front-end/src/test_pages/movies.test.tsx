import {fireEvent, render, screen} from "@testing-library/react";
import Movies from "src/pages/movies";
import userEvent from "@testing-library/user-event";

describe("all movies", () => {
    it("changes the active tab when clicked", () => {
        render(<Movies />);
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

    it("renders Movies component and switches tabs correctly", () => {
        render(<Movies />);

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
