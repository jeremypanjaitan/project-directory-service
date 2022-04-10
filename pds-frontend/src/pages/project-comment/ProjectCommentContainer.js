import ProjectCommentView from "./ProjectCommentView";
import {useCategoryData, useCommentData, useCreateComment, useProfileData, useProjectData} from "./ProjectCommentRepo";
import Swal from "sweetalert2";
import {useLoading} from "../../context";
import {useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import parse from "html-react-parser";
import {useAuth} from "../../services";

const ProjectContainer = () => {
	 const createComment = useCreateComment()
		const commentData = useCommentData()
		const [comments, setComments] = useState([])
		const [shownComments, setShownComments] = useState([])
		const { setIsLoading } = useLoading();
		const params = useParams();
		let [currentPage, setCurrentPage] = useState(1);
		const [copyArray, setCopyArray] = useState([])
		const [totalPages, setTotalPages] = useState(0)
		const projectData = useProjectData()
		const categoryData = useCategoryData()
		const profileData = useProfileData()
		const [project, setProject] = useState();
		const [profile, setProfile] = useState();
		const [category, setCategory] = useState();
		const [tempComment, setTempComment] = useState();
		const auth = useAuth()

		const handleCreateComment = async (dataToPost) => {
				setIsLoading(true);
				try {
						createComment.start({id: params.id, dataToPost: dataToPost});
						setTempComment(dataToPost)
				} catch (err) {
						Swal.fire({
								icon: "error",
								text: "Failed to create comment!",
						});
				} finally {
						setIsLoading(false);
				}
		}

		if (createComment.isSuccess) {
				Swal.fire("comment posted!", "", "success");
				const newest = {picture: auth.user.picture, from: auth.user.fullName, body: tempComment, createdAt: Date.now()}
				setShownComments([newest,...shownComments])
				setIsLoading(false);
				createComment.reset();
		}
		if (createComment.isError) {
				setIsLoading(false);
		}

		useEffect(() => {
				setIsLoading(true)
				projectData.start(params.id);
				commentData.start({id: params.id, page: currentPage})
				//eslint-disable-next-line
		}, []);

		useEffect(() => {
				if (projectData.isSuccess) {
						profileData.start(projectData.data.userId);
						categoryData.start(3);
						projectData.data.story = parse(projectData.data.story);
						setProject(projectData.data);
						projectData.reset();
						setIsLoading(false);
				}

				//eslint-disable-next-line
		}, [projectData.isSuccess, setIsLoading]);

		if (profileData.isSuccess) {
				setProfile(profileData.data);
				profileData.reset();
				setIsLoading(false);
		}

		if (categoryData.isSuccess) {
				setCategory(categoryData.data);
				categoryData.reset();
				setIsLoading(false);
		}

		useEffect(() => {
				if(commentData.isSuccess) {
						setTotalPages(commentData.data.data.totalPage)
						setCurrentPage(commentData.data.data.currentPage)
						const newArray = [...shownComments, ...commentData.data.data.row]
						setShownComments(newArray)
						commentData.reset();
						setIsLoading(false)
				}
		}, [commentData.isSuccess, setIsLoading])

		const handleLoadMore = (val) => {
				setCurrentPage(val)
				setIsLoading(true)
				commentData.start({id: params.id, page: val})
		}

		return (
				<>
						<ProjectCommentView
								profile={profile || {}}
								project={project || {}}
								category={category || {}}
								totalPage={totalPages}
								comments={shownComments || []}
								currentPage={currentPage}
								setCurrentPage={setCurrentPage}
								handleCreateComment={handleCreateComment}
								handleLoadMore={handleLoadMore}
						/>
				</>
		);
};

export default ProjectContainer;
