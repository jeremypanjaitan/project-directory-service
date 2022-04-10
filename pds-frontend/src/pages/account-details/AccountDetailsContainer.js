import { useEffect, useState } from "react";
import { useLoading } from "../../context";
import { useProfileActivity, useProfileProject } from "./AccountDetailsRepo";
import AccountDetailsView from "./AccountDetailsView";

const ProjectContainer = () => {
  const profileProject = useProfileProject();
  const [currentPageProject, setCurrentPageProject] = useState(1);
  const [shownProjects, setShownProjects] = useState([]);
  const [totalPagesProject, setTotalPagesProject] = useState(0);

  const profileActivity = useProfileActivity();
  const [currentPageActivity, setCurrentPageActivity] = useState(1);
  const [shownActivities, setShownActivities] = useState([]);
  const [totalPagesActivity, setTotalPagesActivity] = useState(0);

  const { setIsLoading } = useLoading();

  useEffect(() => {
    profileProject.start({ page: currentPageProject });
    profileActivity.start({ page: currentPageActivity });
    //eslint-disable-next-line
  }, []);

  useEffect(() => {
    if (profileActivity.isSuccess) {
      setTotalPagesActivity(profileActivity.data.data.totalPage);
      setCurrentPageActivity(profileActivity.data.data.currentPage);
      const newArray = [...shownActivities, ...profileActivity.data.data.row];

      setShownActivities(newArray);
      profileActivity.reset();
      setIsLoading(false);
    }

    //eslint-disable-next-line
  }, [profileActivity.isSuccess, setIsLoading]);

  useEffect(() => {
    if (profileProject.isSuccess) {
      setTotalPagesProject(profileProject.data.data.totalPage);
      setCurrentPageProject(profileProject.data.data.currentPage);
      const newArray = [...shownProjects, ...profileProject.data.data.row];

      setShownProjects(newArray);
      profileProject.reset();
      setIsLoading(false);
    }

    //eslint-disable-next-line
  }, [profileProject.isSuccess, setIsLoading]);

  const handleLoadMoreActivity = (val) => {
    setCurrentPageActivity(val);
    setIsLoading(true);
    profileActivity.start({ page: val });
  };
  const handleLoadMoreProject = (val) => {
    setCurrentPageProject(val);
    setIsLoading(true);
    profileProject.start({ page: val });
  };

  return (
    <>
      <AccountDetailsView
        totalPageActivity={totalPagesActivity}
        activities={shownActivities || []}
        currentPageActivity={currentPageActivity}
        setCurrentPageActivity={setCurrentPageActivity}
        handleLoadMoreActivity={handleLoadMoreActivity}
        totalPageProject={totalPagesProject}
        projects={shownProjects || []}
        currentPageProject={currentPageProject}
        setCurrentPageProject={setCurrentPageProject}
        handleLoadMoreProject={handleLoadMoreProject}
      />
    </>
  );
};

export default ProjectContainer;
