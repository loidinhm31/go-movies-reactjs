import React from "react";
import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import EditEpisode from "@/components/Episode/EditEpisode";
import { ManageEpisodesTable } from "@/components/Tables/ManageEpisodesTable";
import ManageMoviesTable from "@/components/Tables/ManageMoviesTable";
import { fakeSearchData } from "@/__tests__/__mocks__/fakeData/search";
import { useRouter } from "next/router";
import { useSession } from "next-auth/react";
import { useDispatch } from "react-redux";
import EditMovie from "@/pages/admin/manage-catalogue/movies";
import EditSeason from "@/pages/admin/manage-catalogue/movies/seasons";


jest.mock("next/router", () => ({
  ...(jest.requireActual("next/router") as object),
  useRouter: jest.fn()
}));

jest.mock("next-auth/react", () => ({
  ...(jest.requireActual("next-auth/react") as object),
  useSession: jest.fn()
}));

jest.mock("react-redux");

describe("ManageMoviesTable", () => {
  test("renders table with movies data", async () => {
    render(<ManageMoviesTable selectedMovie={fakeSearchData.content[0]} setNotifyState={jest.fn}
                              setOpenSeasonDialog={jest.fn} setSelectedMovie={jest.fn} />);

    // Wait for table to load
    await waitFor(() => {
      expect(screen.getByRole("table")).toBeInTheDocument();

    });

    expect(screen.getAllByRole("checkbox")).toHaveLength(7);

    // Assert movie titles
    const movieTitles = screen.getAllByRole("cell", { name: /movie/i });
    expect(movieTitles.length - 1).toBe(fakeSearchData.content.length);
    expect(movieTitles[0]).toHaveTextContent(fakeSearchData.content[0].title);
  });

  test("handles update avg price for tv series", async () => {
    const mockSetNotifyState = jest.fn();

    render(<ManageMoviesTable selectedMovie={fakeSearchData.content[4]} setNotifyState={mockSetNotifyState}
                              setOpenSeasonDialog={jest.fn} setSelectedMovie={jest.fn} />);

    // Wait for table to load
    await waitFor(() => {
      expect(screen.getByRole("table")).toBeInTheDocument();
    });

    // Click delete button
    const updateAvgPriceButton = screen.getByTestId(`update-button-${fakeSearchData.content[4].id}`);
    userEvent.click(updateAvgPriceButton);

    await waitFor(() => {
      expect(mockSetNotifyState).toBeCalled();
    });
  });

  test("handles view seasons for tv series", async () => {
    const mockOpenSeasonDialog = jest.fn();

    render(<ManageMoviesTable selectedMovie={fakeSearchData.content[4]} setNotifyState={jest.fn}
                              setOpenSeasonDialog={mockOpenSeasonDialog} setSelectedMovie={jest.fn} />);

    // Wait for table to load
    await waitFor(() => {
      expect(screen.getByRole("table")).toBeInTheDocument();
    });

    // Click delete button
    const viewSeasonButton = screen.getByTestId(`view-button-${fakeSearchData.content[4].id}`);
    userEvent.click(viewSeasonButton);

    await waitFor(() => {
      expect(mockOpenSeasonDialog).toBeCalled();
    });
  });

  test("handles movie deletion", async () => {
    const mockSetNotifyState = jest.fn();
    render(<ManageMoviesTable selectedMovie={fakeSearchData.content[0]} setNotifyState={mockSetNotifyState}
                              setOpenSeasonDialog={jest.fn} setSelectedMovie={jest.fn} />);

    // Wait for table to load
    await waitFor(() => {
      expect(screen.getByRole("table")).toBeInTheDocument();
    });

    // Click delete button
    const deleteButton = screen.getByTestId(`del-button-${fakeSearchData.content[0].id}`);
    userEvent.click(deleteButton);

    // Confirm deletion
    const confirmButton = screen.getByRole("button", { name: /yes/i });
    userEvent.click(confirmButton);

    await waitFor(() => {
      expect(mockSetNotifyState).toBeCalled();

      // Check page was refreshed with removed item
      expect(screen.getByRole("table")).toBeInTheDocument();
    });
  });
});

describe("Edit Movie", () => {
  test("renders edit movie form", async () => {
    (useRouter as jest.Mock).mockImplementation(() => ({
      query: {
        id: "1"
      },
      push: jest.fn()
    }));

    (useSession as jest.Mock).mockImplementation(() => ({
      data: {
        expires: new Date(Date.now() + 2 * 86400).toISOString(),
        user: { username: "admin" }
      },
      status: "authenticated"
    }));

    // @ts-ignore
    render(<EditMovie />);

    await waitFor(() => {
      expect(screen.getByLabelText("Title")).toHaveDisplayValue("Test movie 1");
      expect(screen.getByLabelText("Price")).toHaveDisplayValue("0");
      expect(screen.getByLabelText("Release Date")).toHaveDisplayValue("2014-06-05");
      expect(screen.getByLabelText("Runtime")).toHaveDisplayValue("102");
      expect(screen.getByLabelText("Description")).toHaveTextContent("Test desc movie 1");
      expect(screen.getByRole("checkbox", { name: "G1" })).toBeChecked();
      expect(screen.getByRole("checkbox", { name: "G2" })).toBeChecked();
    });
  });

  test("handles form submission", async () => {
    const mockPush = jest.fn();
    (useRouter as jest.Mock).mockImplementation(() => ({
      query: {},
      push: mockPush
    }));

    (useSession as jest.Mock).mockImplementation(() => ({
      data: {
        expires: new Date(Date.now() + 2 * 86400).toISOString(),
        user: { username: "admin" }
      },
      status: "authenticated"
    }));

    (useDispatch as jest.Mock).mockReturnValue(jest.fn());

    render(<EditMovie />);

    await waitFor(async () => {
      expect(screen.getByLabelText("MPAA Rating")).toBeInTheDocument();

      expect(screen.getByText("Genres")).toBeInTheDocument();
    });

    // Select Movie Type
    userEvent.click(screen.getByRole("radio", { name: "Movie" }));

    fireEvent.change(screen.getByLabelText("Title"), { target: { value: "Test Movie" } });

    fireEvent.change(screen.getByLabelText("Price"), { target: { value: 9.99 } });

    fireEvent.change(screen.getByLabelText("Description"), { target: { value: "Test Desc Movie" } });

    userEvent.click(screen.getByLabelText("MPAA Rating"));

    userEvent.click(screen.getByRole("option", { name: "PG" }));

    // Check the action genre checkbox
    await waitFor(() => {
      userEvent.click(screen.getByRole("checkbox", { name: "G1" }));
      userEvent.click(screen.getByRole("checkbox", { name: "G2" }));
    });

    // Submit the form
    userEvent.click(screen.getByText("Save"));

    await waitFor(() => {
      expect(mockPush).toBeCalledWith("/admin/manage-catalogue");
    });
  });

  test("handles submission error", async () => {
    (useRouter as jest.Mock).mockImplementation(() => ({
      query: {},
      push: jest.fn()
    }));

    (useSession as jest.Mock).mockImplementation(() => ({
      data: {
        expires: new Date(Date.now() + 2 * 86400).toISOString(),
        user: { username: "admin" }
      },
      status: "authenticated"
    }));
    render(<EditMovie />);

    // Select Movie Type
    userEvent.click(screen.getByRole("radio", { name: "TV Series" }));

    fireEvent.change(screen.getByLabelText("Title"), { target: { value: "Test TV series" } });

    fireEvent.change(screen.getByLabelText("Description"), { target: { value: "Test Desc TV series" } });

    // Submit the form
    userEvent.click(screen.getByText("Save"));

    await waitFor(() => {
      expect(screen.getByText("Fill value for MPAA Rating, Genres")).toBeInTheDocument();

      expect(screen.getByText("You must choose at least one genre!")).toBeInTheDocument();
    });
  });

  test("handles delete movie", async () => {
    const mockPush = jest.fn();

    (useRouter as jest.Mock).mockImplementation(() => ({
      query: {
        id: "1"
      },
      push: mockPush
    }));

    (useSession as jest.Mock).mockImplementation(() => ({
      data: {
        expires: new Date(Date.now() + 2 * 86400).toISOString(),
        user: { username: "admin" }
      },
      status: "authenticated"
    }));

    (useDispatch as jest.Mock).mockReturnValue(jest.fn());

    render(<EditMovie />);

    await waitFor(async () => {
      expect(screen.getByRole("button", { name: "Save" })).toBeInTheDocument();
      expect(screen.getByText("Delete Movie")).toBeInTheDocument();
    });

    userEvent.click(screen.getByText("Delete Movie"));

    expect(screen.getByText("Delete Item")).toBeInTheDocument();

    userEvent.click(screen.getByText("Yes"));

    await waitFor(() => {
      expect(mockPush).toBeCalledWith("/admin/manage-catalogue");

    });
  });

  test("handles file video upload", async () => {
    (useRouter as jest.Mock).mockImplementation(() => ({
      query: {},
      push: jest.fn()
    }));

    (useSession as jest.Mock).mockImplementation(() => ({
      data: {
        expires: new Date(Date.now() + 2 * 86400).toISOString(),
        user: { username: "admin" }
      },
      status: "authenticated"
    }));

    render(<EditMovie />);

    // Select Movie Type
    userEvent.click(screen.getByRole("radio", { name: "Movie" }));

    // Mock the file input change event
    const videoFile = new File(["dummy video content"], "example-file.mp4", { type: "video/mp4" });
    const fileInput = screen.getByTestId("uploadVideo");
    fireEvent.change(fileInput, { target: { files: [videoFile] } });

    await waitFor(() => {
      // Check that the file was uploaded and the state was updated
      expect(screen.getByText("example-file")).toBeInTheDocument();
      expect(screen.getByText("Video Uploaded")).toBeInTheDocument();
    });

  });

  test("handles file image upload", async () => {
    (useRouter as jest.Mock).mockImplementation(() => ({
      query: {},
      push: jest.fn()
    }));

    (useSession as jest.Mock).mockImplementation(() => ({
      data: {
        expires: new Date(Date.now() + 2 * 86400).toISOString(),
        user: { username: "admin" }
      },
      status: "authenticated"
    }));

    render(<EditMovie />);

    // Select Movie Type
    userEvent.click(screen.getByRole("radio", { name: "Movie" }));

    // Mock the file input change event
    const imageFile = new File(["dummy image"], "example-file.jpg", { type: "image/jpeg" });
    const fileInput = screen.getByTestId("uploadImage");
    fireEvent.change(fileInput, { target: { files: [imageFile] } });

    await waitFor(() => {
      // Check that the file was uploaded and the state was updated
      expect(screen.getByLabelText("Image Path")).toHaveValue("example-file");
      expect(screen.getByText("Image Uploaded")).toBeInTheDocument();
    });

  });

  test("Update TV series Avg price", async () => {
    (useRouter as jest.Mock).mockImplementation(() => ({
      query: {
        id: "5"
      },
      push: jest.fn()
    }));

    (useSession as jest.Mock).mockImplementation(() => ({
      data: {
        expires: new Date(Date.now() + 2 * 86400).toISOString(),
        user: { username: "admin" }
      },
      status: "authenticated"
    }));

    render(<EditMovie />);

    await waitFor(async () => {
      userEvent.click(screen.getByTestId("updateAvgPrice"));

      expect(await screen.getByText("Average Price Was Updated")).toBeInTheDocument();

    });

  });
});

describe("Edit Season", () => {
  test("should render the component and data", async () => {
    (useRouter as jest.Mock).mockImplementation(() => ({
      query: {
        id: "1",
        movieId: "3"
      },
      push: jest.fn()
    }));

    (useSession as jest.Mock).mockImplementation(() => ({
      data: {
        expires: new Date(Date.now() + 2 * 86400).toISOString(),
        user: { username: "admin" }
      },
      status: "authenticated"
    }));

    render(<EditSeason />);

    // Assert that the component is rendered correctly
    expect(screen.getByText("Add/Edit Season")).toBeInTheDocument();
    expect(screen.getByLabelText("Name")).toBeInTheDocument();
    expect(screen.getByLabelText("Release Date")).toBeInTheDocument();
    expect(screen.getByLabelText("Description")).toBeInTheDocument();
    expect(screen.getByText("Save")).toBeInTheDocument();

    await waitFor(() => {
      expect(screen.getByLabelText("Name")).toHaveDisplayValue("Season 1");
      expect(screen.getByLabelText("Release Date")).toHaveDisplayValue("2019-06-01");
      expect(screen.getByLabelText("Description")).toHaveTextContent("First Season at OT");
    });
  });

  test("should display an error message when failed to fetch season data", async () => {
    (useRouter as jest.Mock).mockImplementation(() => ({
      query: {
        id: "2",
        movieId: "3"
      },
      push: jest.fn()
    }));

    (useSession as jest.Mock).mockImplementation(() => ({
      data: {
        expires: new Date(Date.now() + 2 * 86400).toISOString(),
        user: { username: "admin" }
      },
      status: "authenticated"
    }));

    render(<EditSeason />);

    // Assert that an error message is displayed when fetching season data fails
    await waitFor(() => {
      expect(screen.getByText("server error")).toBeInTheDocument();
    });
  });
//
  test("should save season data when 'Save' button is clicked", async () => {
    const mockPush = jest.fn();
    (useRouter as jest.Mock).mockImplementation(() => ({
      query: {},
      push: mockPush
    }));


    (useSession as jest.Mock).mockImplementation(() => ({
      data: {
        expires: new Date(Date.now() + 2 * 86400).toISOString(),
        user: { username: "admin" }
      },
      status: "authenticated"
    }));

    render(<EditSeason />);

    // // Fill in the form fields
    fireEvent.change(screen.getByLabelText("Name"), { target: { value: "Season 1" } });
    fireEvent.change(screen.getByLabelText("Release Date"), { target: { value: "2023-01-01" } });
    fireEvent.change(screen.getByLabelText("Description"), { target: { value: "Test season" } });

    // Click the 'Save' button
    fireEvent.click(screen.getByText("Save"));

    // Assert that the success message is displayed
    await waitFor(async () => {
      expect(mockPush).toHaveBeenCalledWith("/admin/manage-catalogue");
    });
  });

  test("should have an error when 'Save' button is clicked", async () => {
    const mockPush = jest.fn();
    (useRouter as jest.Mock).mockImplementation(() => ({
      query: {},
      push: mockPush
    }));


    (useSession as jest.Mock).mockImplementation(() => ({
      data: {
        expires: new Date(Date.now() + 2 * 86400).toISOString(),
        user: { username: "admin" }
      },
      status: "authenticated"
    }));

    render(<EditSeason />);

    // // Fill in the form fields
    fireEvent.change(screen.getByLabelText("Release Date"), { target: { value: "2023-01-01" } });

    // Click the 'Save' button
    fireEvent.click(screen.getByText("Save"));

    expect(screen.getByText("Fill value for Name, Description")).toBeInTheDocument();
  });

  test("should delete season when 'Delete Season' button is clicked", async () => {
    const mockPush = jest.fn();
    (useRouter as jest.Mock).mockImplementation(() => ({
      query: {
        id: "1",
        movieId: "3"
      },
      push: mockPush
    }));

    (useSession as jest.Mock).mockImplementation(() => ({
      data: {
        expires: new Date(Date.now() + 2 * 86400).toISOString(),
        user: { username: "admin" }
      },
      status: "authenticated"
    }));

    render(<EditSeason />);

    await waitFor(() => {
      expect(screen.getByLabelText("Name")).toHaveDisplayValue("Season 1");
    });

    // Click the 'Delete Season' button
    fireEvent.click(screen.getByText("Delete Season"));

    // Assert that the confirmation dialog is displayed
    expect(screen.getByText("Delete Item")).toBeInTheDocument();

    // Confirm the deletion
    fireEvent.click(screen.getByText("Yes"));

    await waitFor(() => {
      expect(mockPush).toBeCalledWith("/admin/manage-catalogue");
    });

  });
});

describe("Edit Episode", () => {
  test("renders ManageEpisodesTable component", async () => {
    render(
      <ManageEpisodesTable
        season={{
          id: 5,
          name: "Season 5",
          air_date: "2022-06-01T00:00:00Z",
          description: "Fifth Season at OT",
          movie_id: 5
        }}
        setNotifyState={jest.fn} />
    );

    await waitFor(() => {
      expect(screen.getByText("E1: Test episode 1")).toBeInTheDocument();
      expect(screen.getByText("E2: Test episode 2")).toBeInTheDocument();
    });
  });

  test("should handle form submission", async () => {
    const mockSetNotifyState = jest.fn();
    const mockSetWasUpdated = jest.fn();

    render(<EditEpisode id={1} seasonId={1} setNotifyState={mockSetNotifyState} setWasUpdated={mockSetWasUpdated} />);

    // Open edit
    userEvent.click(screen.getByTestId("edit-episode"));

    // Fill in the form fields
    const nameField = screen.getByLabelText("Name");
    userEvent.click(nameField);
    userEvent.type(nameField, "Change Name");

    const runtimeField = screen.getByLabelText("Runtime");
    userEvent.click(runtimeField);
    userEvent.type(runtimeField, "10");

    fireEvent.change(screen.getByLabelText("Air Date"), { target: { value: "2023-06-30" } });

    // Submit the form
    userEvent.click(screen.getByRole("button", { name: "Save" }));

    // Wait for the form submission to complete
    await waitFor(() => {
      expect(mockSetNotifyState).toBeCalledWith({
        open: true,
        message: "Episode Saved",
        vertical: "top",
        horizontal: "right",
        severity: "success"
      });

      expect(mockSetWasUpdated).toBeCalledWith(true);
    });
  });

  test("Delete episode", async () => {
    const mockSetNotifyState = jest.fn();
    const mockSetWasUpdated = jest.fn();

    render(<EditEpisode id={1} seasonId={1} setNotifyState={mockSetNotifyState} setWasUpdated={mockSetWasUpdated} />);

    userEvent.click(screen.getByTestId("delete-episode"));

    userEvent.click(screen.getByText("Yes"));

    await waitFor(() => {
      expect(mockSetNotifyState).toBeCalledWith({
        open: true,
        message: "ok",
        vertical: "top",
        horizontal: "right",
        severity: "info"
      });

      expect(mockSetWasUpdated).toBeCalledWith(true);
    });
  });

  test("Missing field when adding episode", async () => {
    const mockSetNotifyState = jest.fn();

    render(<EditEpisode seasonId={1} setNotifyState={mockSetNotifyState} setWasUpdated={jest.fn} />);

    userEvent.click(screen.getByTestId("add-episode"));

    const nameInput = screen.getByLabelText("Name");
    userEvent.click(nameInput);
    userEvent.type(nameInput, "test");

    userEvent.click(screen.getByRole("button", { name: "Save" }));

    expect(mockSetNotifyState).toBeCalledWith({
      open: true,
      message: `Fill value for Runtime`,
      vertical: "bottom",
      horizontal: "center",
      severity: "warning"
    });
  });

  test("handles file video upload", async () => {
    const mockSetNotifyState = jest.fn();

    render(<EditEpisode seasonId={1} setNotifyState={mockSetNotifyState} setWasUpdated={jest.fn} />);

    userEvent.click(screen.getByTestId("add-episode"));

    // Mock the file input change event
    const videoFile = new File(["dummy video content"], "example-file.mp4", { type: "video/mp4" });
    const fileInput = screen.getByTestId("upload-video");
    fireEvent.change(fileInput, { target: { files: [videoFile] } });

    await waitFor(() => {
      // Check that the file was uploaded and the state was updated
      expect(screen.getByText("example-file")).toBeInTheDocument();
      expect(mockSetNotifyState).toBeCalledWith({
        open: true,
        message: "Video Uploaded",
        vertical: "top",
        horizontal: "right",
        severity: "info",
      })
    });

  });
});